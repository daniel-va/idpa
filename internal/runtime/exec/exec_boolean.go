package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecBoolean(ctx Context, node ast.BooleanNode) (Value, *Error) {
    return BooleanValue(node.Value), nil
}
