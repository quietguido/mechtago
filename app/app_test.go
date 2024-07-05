package app

import (
	"runtime"
	"testing"

	"github.com/quietguido/mechtago/calculator"
	"github.com/quietguido/mechtago/generator"
)

const (
	inputFileName = "testing"
	itemCount     = 1e6
	workerCount   = 4
)

func BenchmarkCalculate(b *testing.B) {
	groundTruth := generator.GenerateJson(inputFileName, itemCount)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		answ := calculator.Calculate(inputFileName, runtime.NumCPU())

		if groundTruth != answ {
			b.FailNow()
		}
	}
}
