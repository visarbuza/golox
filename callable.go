package main

type Callable interface {
	Arity() int
	Call(i *Interpreter, arguments []any) any
}

type LoxFunction struct {
	Closure     *Environment
	Declaration *FunctionStmt
}

func NewLoxFunction(declaration *FunctionStmt, closure *Environment) *LoxFunction {
	return &LoxFunction{
		Declaration: declaration,
		Closure:     closure,
	}
}

func (f *LoxFunction) Arity() int {
	return len(f.Declaration.Params)
}

func (f *LoxFunction) Call(i *Interpreter, arguments []any) (returned any) {
	defer func() {
		if r := recover(); r != nil {
			if signal, ok := r.(ReturnSignal); ok {
				returned = signal.Value
			} else {
				panic(r)
			}
		}
	}()
	env := NewEnvironment(f.Closure)
	for i, param := range f.Declaration.Params {
		env.Define(param.Lexeme, arguments[i])
	}
	i.executeBlock(f.Declaration.Body, env)
	return nil
}

func (f *LoxFunction) String() string {
	return "<fn " + f.Declaration.Name.Lexeme + ">"
}
