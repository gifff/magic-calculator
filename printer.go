package main

import (
	"fmt"
	"os"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error occured when executing program: %s\n", err)
	os.Exit(1)
}
