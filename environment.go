package main

import "fmt"

var values = map[string]any{}

func define(name string, value any) {
	values[name] = value
}

func get(name Token) (any, error) {
	if val, ok := values[name.Lexeme]; ok {
		return val, nil
	}
	return nil, &RuntimeError{Token: name, Message: fmt.Sprintf("Undefined variable %s.", name.Lexeme)}
}
