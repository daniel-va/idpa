package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecInteger(ctx Context, node ast.IntegerNode) (Value, *Error) {
    return NumberValue(float64(node.Value)), nil
}

