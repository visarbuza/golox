package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError = false

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox[script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading input:", err)
			continue
		}
		// Remove the newline character
		line = line[:len(line)-1]
		run(line)
		hadError = false
	}

	return nil
}

func run(source string) {
	scanner := NewScanner(source)

	for _, token := range scanner.ScanTokens() {
		fmt.Println(token)
	}
}

func err(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}
