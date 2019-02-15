package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecString(ctx Context, node ast.StringNode) (Value, *Error) {
    return StringValue(node.UsableValue()), nil
}