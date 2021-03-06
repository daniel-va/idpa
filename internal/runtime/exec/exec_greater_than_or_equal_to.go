package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecGreaterThanOrEqualTo(ctx Context, node ast.GreaterThanOrEqualToNode) (Value, *Error) {
    return execInfix(ctx, node, InfixMapping{
        Message: "can't compare %s to %s",
        Number: func(a, b float64) (Value, *Error) {
            return BooleanValue(a >= b), nil
        },
    })
}