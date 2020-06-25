package main

import (
	"fmt"
	"strings"
)

func FormatSequenceResult(r SequenceResult) string {
	sb := &strings.Builder{}
	for i, rr := range r {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", rr))
	}

	return sb.String()
}
