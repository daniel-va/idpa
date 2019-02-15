package library

import (
    "fmt"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func initIO() {
    Entries["print"] = ClosureValue{
        Arity: 1,
        Call: func(values []Value) (Value, *Error) {
            fmt.Print(values[0].Dump())
            return NullValue{}, nil
        },
    }

    Entries["println"] = ClosureValue{
        Arity: 1,
        Call: func(values []Value) (Value, *Error) {
            fmt.Println(values[0].Dump())
            return NullValue{}, nil
        },
    }
}