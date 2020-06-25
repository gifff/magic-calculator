package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciEvaluator(t *testing.T) {
	t.Run("n=0 result should be empty", func(t *testing.T) {
		sb := &strings.Builder{}
		e := &FirstFibonacciEvaluator{N: 0, ResultWriter: sb}
		var evaluator Evaluator = e
		evaluator.Evaluate()

		assert.NotNil(t, e.Result)
		assert.ElementsMatch(t, []int64{}, e.Result)
		assert.Equal(t, "Result: \n", sb.String())
	})
}
