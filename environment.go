package main

import "fmt"

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing: enclosing, values: make(map[string]any)}
}

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Get(name Token) (any, error) {
	if val, ok := e.values[name.Lexeme]; ok {
		return val, nil
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	return nil, &RuntimeError{Token: name, Message: fmt.Sprintf("Undefined variable %s.", name.Lexeme)}
}

func (e *Environment) Assign(name Token, value any) {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return
	}
	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
		return
	}
	panic(&RuntimeError{Token: name, Message: fmt.Sprintf("Undefined variable %s.", name.Lexeme)})
}
