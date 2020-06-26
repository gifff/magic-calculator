package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	var (
		args []string
		cmd  string
		err  error
	)

	printHelp(true)
	fmt.Printf("> ")

	for scanner.Scan() {
		args = strings.Split(scanner.Text(), " ")
		if len(args) == 0 {
			fmt.Printf("> ")
			continue
		}

		cmd = args[0]
		args = args[1:]

		switch cmd {
		case "help":
			printHelp(true)
			fmt.Printf("> ")
			continue
		case "quit", "exit":
			return
		}

		if !checkRequiredArgs(cmd, args) {
			fmt.Printf("> ")
			continue
		}

		err = evaluate(cmd, args)
		if err != nil {
			fmt.Printf("Error occured: %s\n", err)
		}

		fmt.Printf("> ")
	}

	err = scanner.Err()
	if err != nil {
		printError(err)
	}
}
