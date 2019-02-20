package library

import (
    "bufio"
    "fmt"
    . "github.com/daniel-va/idpa/internal/runtime"
    "os"
    "strings"
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

    Entries["read_line"] = ClosureValue{
        Arity: 0,
        Call: func(values []Value) (Value, *Error) {
            reader := bufio.NewReader(os.Stdin)
            text, err := reader.ReadString('\n')
            if err != nil {
                return nil, Err("failed to read from console: %s", err)
            }
            text = strings.TrimSuffix(text, "\n")
            return StringValue(text), nil
        },
    }
}