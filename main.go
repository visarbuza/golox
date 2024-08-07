package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError = false
var hadRuntimeError = false
var interpreter = &Interpreter{env: NewEnvironment(nil)}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: golox[script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
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
	if hadRuntimeError {
		os.Exit(70)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()
	scanner := NewScanner(source)
	parser := NewParser(scanner.ScanTokens())

	statements := parser.Parse()

	if hadError {
		return
	}

	interpreter.Interpret(statements)
}

func err(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}

func errLox(token Token, message string) {
	if token.Type == EOF {
		report(token.Line, " at end", message)
	} else {
		report(token.Line, fmt.Sprintf(" at '%s'", token.Lexeme), message)
	}
}

func runtimeError(err RuntimeError) {
	fmt.Printf("%s\n[line %d]\n", err.Message, err.Token.Line)
	hadRuntimeError = true
}
