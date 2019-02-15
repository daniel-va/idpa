package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecFloat(ctx Context, node ast.FloatNode) (Value, *Error) {
    return NumberValue(node.Value), nil
}
