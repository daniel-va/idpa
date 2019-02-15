package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecMultiplication(ctx Context, node ast.MultiplicationNode) (Value, *Error) {
    return execInfix(ctx, node, InfixMapping{
        Message: "can't multiply %s by %s",
        Number: func(a, b float64) (Value, *Error) {
            return NumberValue(a * b), nil
        },
    })
}