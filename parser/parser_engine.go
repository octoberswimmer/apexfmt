package parser

import (
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// maxParserEngines bounds the number of parser ATN copies. Each copy builds
// its own DFA cache, so more copies mean more CPU spent on prediction in a
// fresh process; four copies gave the best wall time on the afmt benchmark
// corpus with 28 threads.
const maxParserEngines = 4

// ParserEngine holds a copy of the parser ATN together with the DFA cache and
// prediction context cache built from it. Parsers sharing one ATN across
// goroutines contend on its DFA edge lock while the cache is being built, so
// concurrent parsers are spread over a few copies. An engine may be used by
// several parsers at once; the antlr runtime guards its state.
type ParserEngine struct {
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
	predictionContextCache *antlr.PredictionContextCache
	users                  int
}

var (
	parserEnginesMu sync.Mutex
	parserEngines   []*ParserEngine
)

func newParserEngine() *ParserEngine {
	ApexParserInit()
	atn := antlr.NewATNDeserializer(nil).Deserialize(ApexParserParserStaticData.serializedATN)
	decisionToDFA := make([]*antlr.DFA, len(atn.DecisionToState))
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
	return &ParserEngine{
		atn:                    atn,
		decisionToDFA:          decisionToDFA,
		predictionContextCache: antlr.NewPredictionContextCache(),
	}
}

// AcquireParserEngine returns the engine with the fewest parsers in use,
// creating a new one when every existing engine is busy and fewer than
// maxParserEngines exist. A single goroutine parsing files one after another
// therefore always gets the first engine, whose DFA cache is warm. The caller
// must call release when done with the engine's parsers.
func AcquireParserEngine() (engine *ParserEngine, release func()) {
	parserEnginesMu.Lock()
	defer parserEnginesMu.Unlock()
	for _, e := range parserEngines {
		if engine == nil || e.users < engine.users {
			engine = e
		}
	}
	if (engine == nil || engine.users > 0) && len(parserEngines) < maxParserEngines {
		engine = newParserEngine()
		parserEngines = append(parserEngines, engine)
	}
	engine.users++
	return engine, func() {
		parserEnginesMu.Lock()
		engine.users--
		parserEnginesMu.Unlock()
	}
}

// NewApexParser produces a parser that predicts with this engine's ATN copy.
func (e *ParserEngine) NewApexParser(input antlr.TokenStream) *ApexParser {
	p := NewApexParser(input)
	p.Interpreter = antlr.NewParserATNSimulator(p, e.atn, e.decisionToDFA, e.predictionContextCache)
	return p
}
