package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecNameReference(ctx Context, node ast.NameReferenceNode) (Value, *Error) {
    value, ok := ctx.Fetch(node.Name)
    if !ok {
        return nil, Err("unknown identifier `%s`", node.Name).AtNode(node)
    }
    return value, nil
}