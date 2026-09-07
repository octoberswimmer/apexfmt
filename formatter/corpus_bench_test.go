package formatter

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// BenchmarkCorpus formats every .cls and .trigger file under the directory
// named by APEXFMT_BENCH_DIR. It is skipped when the variable is unset so the
// regular test run does not depend on an external corpus.
//
//	APEXFMT_BENCH_DIR=/path/to/corpus go test ./formatter -run '^$' -bench Corpus -cpuprofile cpu.out
func BenchmarkCorpus(b *testing.B) {
	dir := os.Getenv("APEXFMT_BENCH_DIR")
	if dir == "" {
		b.Skip("APEXFMT_BENCH_DIR not set")
	}
	sources := readCorpus(b, dir)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, src := range sources {
			f := NewFormatter("", bytes.NewReader(src))
			if err := f.Format(); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkCorpusParallel formats the same files concurrently, one goroutine
// per file with GOMAXPROCS of them running at once, the way the command line
// does with --write or --list.
func BenchmarkCorpusParallel(b *testing.B) {
	dir := os.Getenv("APEXFMT_BENCH_DIR")
	if dir == "" {
		b.Skip("APEXFMT_BENCH_DIR not set")
	}
	sources := readCorpus(b, dir)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sem := make(chan struct{}, runtime.GOMAXPROCS(0))
		var wg sync.WaitGroup
		for _, src := range sources {
			wg.Add(1)
			go func(src []byte) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				f := NewFormatter("", bytes.NewReader(src))
				if err := f.Format(); err != nil {
					b.Error(err)
				}
			}(src)
		}
		wg.Wait()
	}
}

func readCorpus(b *testing.B, dir string) [][]byte {
	var sources [][]byte
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !(strings.HasSuffix(path, ".cls") || strings.HasSuffix(path, ".trigger")) {
			return nil
		}
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		sources = append(sources, src)
		return nil
	})
	if err != nil {
		b.Fatal(err)
	}
	return sources
}
