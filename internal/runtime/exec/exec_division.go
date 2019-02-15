package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecDivision(ctx Context, node ast.DivisionNode) (Value, *Error) {
    return execInfix(ctx, node, InfixMapping{
        Message: "can't divide %s by %s",
        Number: func(a, b float64) (Value, *Error) {
            if b == 0 {
                return nil, Err("can't divide by 0").AtNode(node)
            }
            return NumberValue(a / b), nil
        },
    })
}