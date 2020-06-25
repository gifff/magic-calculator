package main

import (
	"flag"
	"os"
	"strconv"
)

func main() {
	cmd := flag.String("c", "", "Use specific command. Available commands are: sum")
	flag.Parse()

	input := make([]int64, 0, flag.NArg())
	for _, arg := range flag.Args() {
		v, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			panic(err)
		}

		input = append(input, v)
	}

	if *cmd == "sum" {
		evaluator := &SumEvaluator{
			X:            input[0],
			Y:            input[1],
			ResultWriter: os.Stdout,
		}
		evaluator.Evaluate()
	}
}
