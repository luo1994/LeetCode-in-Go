package problem0842

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []int
}{

	{
		"123456579",
		[]int{123, 456, 579},
	},

	{
		"11235813",
		[]int{1, 1, 2, 3, 5, 8, 13},
	},

	{
		"112358130",
		[]int{},
	},

	{
		"0123",
		[]int{},
	},

	{
		"1101111",
		[]int{110, 1, 111},
	},

	// 可以有多个 testcase
}

func Test_splitIntoFibonacci(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, splitIntoFibonacci(tc.S), "输入:%v", tc)
	}
}

func Benchmark_splitIntoFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			splitIntoFibonacci(tc.S)
		}
	}
}