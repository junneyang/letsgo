package test

import (
	// "fmt"
	"testing"
)

func Benchmark_FlowControl(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		b.Log(i)
	}
}
