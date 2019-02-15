package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecGroup(ctx Context, node ast.GroupNode) (Value, *Error) {
    return exec(ctx, node.Node)
}
