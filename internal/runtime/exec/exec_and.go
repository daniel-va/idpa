package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecAnd(ctx Context, node ast.AndNode) (Value, *Error) {
    return execInfixBoolean(ctx, node, func(a bool, execB func() (bool, *Error)) (bool, *Error) {
        if !a {
            return false, nil
        }
        b, err := execB()
        if err != nil {
            return false, err
        }
        return b, nil
    })
}