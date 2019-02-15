package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecAddition(ctx Context, node ast.AdditionNode) (Value, *Error) {
    return execInfix(ctx, node, InfixMapping{
        Message: "can't add %s to %s",
        Number: func(a, b float64) (Value, *Error) {
            return NumberValue(a + b), nil
        },
        String: func(a, b string) (Value, *Error) {
            return StringValue(a + b), nil
        },
    })
}