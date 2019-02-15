package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecNotEquals(ctx Context, node ast.NotEqualsNode) (Value, *Error) {
    value, err := ExecEquals(ctx, node)
    if err != nil {
        return nil, err
    }
    return !value, nil
}