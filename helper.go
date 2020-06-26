package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func printUsage() {
	executableName := filepath.Base(os.Args[0])
	fmt.Printf("For usage information: %s help\n", executableName)
}

func printHelp(isInteractive bool) {
	executableName := ""
	if !isInteractive {
		executableName = filepath.Base(os.Args[0])
	}

	fmt.Printf("Usage: %s <command> <arg1> <arg2> ... <argN>\n\n", executableName)
	fmt.Printf("Available commands:\n")
	fmt.Printf("\thelp\t\tPrint this help guide\n")
	fmt.Printf("\tsum\t\tSum of <arg1> and <arg2>\n")
	fmt.Printf("\tmultiply\tMultiply <arg1> and <arg2>\n")
	fmt.Printf("\tnfib\t\tPrint the first N Fibonacci number\n")
	fmt.Printf("\tnprime\t\tPrint the first N Prime number\n")
	if isInteractive {
		fmt.Printf("\tquit, exit\tExit the program\n")
	} else {
		fmt.Printf("\tinteractive\tRun program in interactive mode\n")
	}
}
