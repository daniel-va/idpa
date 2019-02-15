package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecSubtraction(ctx Context, node ast.SubtractionNode) (Value, *Error) {
    return execInfix(ctx, node, InfixMapping{
        Message: "can't subtract %[2]s from %[1]s",
        Number: func(a, b float64) (Value, *Error) {
            return NumberValue(a - b), nil
        },
    })
}