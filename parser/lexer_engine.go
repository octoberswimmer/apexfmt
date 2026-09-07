package parser

import (
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// lexerEngine holds a private copy of the lexer ATN together with the DFA
// cache built from it. The antlr runtime guards DFA edge lookups with a
// read-write mutex on the ATN and takes it for every input character, so
// lexers that share one ATN across goroutines contend on that lock. An engine
// is used by one lexer at a time.
type lexerEngine struct {
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
	predictionContextCache *antlr.PredictionContextCache
}

func newLexerEngine() *lexerEngine {
	ApexLexerInit()
	atn := antlr.NewATNDeserializer(nil).Deserialize(ApexLexerLexerStaticData.serializedATN)
	decisionToDFA := make([]*antlr.DFA, len(atn.DecisionToState))
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
	return &lexerEngine{
		atn:                    atn,
		decisionToDFA:          decisionToDFA,
		predictionContextCache: antlr.NewPredictionContextCache(),
	}
}

var lexerEngines = sync.Pool{
	New: func() any { return newLexerEngine() },
}

// NewApexLexerWithPrivateATN produces a lexer whose ATN and DFA cache are not
// shared with lexers running on other goroutines. The caller must call
// release once the lexer is no longer used, which returns the engine to a pool
// so its DFA cache is reused.
func NewApexLexerWithPrivateATN(input antlr.CharStream) (lexer *ApexLexer, release func()) {
	l := NewApexLexer(input)
	engine := lexerEngines.Get().(*lexerEngine)
	l.Interpreter = antlr.NewLexerATNSimulator(l, engine.atn, engine.decisionToDFA, engine.predictionContextCache)
	return l, func() { lexerEngines.Put(engine) }
}
