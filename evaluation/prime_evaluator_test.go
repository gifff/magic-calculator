package evaluation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPrimeEvaluator(t *testing.T) {
	testCases := []struct {
		name                  string
		n                     int64
		expectedResult        SequenceResult
		expectedWrittenResult string
		expectedError         error
	}{
		{
			name:                  "n=0 result should be empty",
			n:                     0,
			expectedResult:        SequenceResult{},
			expectedWrittenResult: "\n",
			expectedError:         nil,
		},
		{
			name:                  "n=1 result should be '2'",
			n:                     1,
			expectedResult:        SequenceResult{2},
			expectedWrittenResult: "2\n",
			expectedError:         nil,
		},
		{
			name:                  "n=10 result should be '2, 3, 5, 7, 11, 13, 17, 19, 23, 29'",
			n:                     10,
			expectedResult:        SequenceResult{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
			expectedWrittenResult: "2, 3, 5, 7, 11, 13, 17, 19, 23, 29\n",
			expectedError:         nil,
		},
		{
			name:                  "n=-1 should return error",
			n:                     -1,
			expectedResult:        SequenceResult{},
			expectedWrittenResult: "",
			expectedError:         ErrInvalidInput,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			e := &FirstPrimeEvaluator{N: tc.n, ResultWriter: sb}
			var evaluator Evaluator = e
			err := evaluator.Evaluate()

			assert.Equal(t, tc.expectedError, err)
			assert.NotNil(t, e.Result)
			assert.ElementsMatch(t, tc.expectedResult, e.Result)
			assert.Equal(t, tc.expectedWrittenResult, sb.String())
		})
	}
}
