package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveName(ctx Context) (node ast.NameNode, success bool) {
    nameTk, ok := ctx.Expect(token.Kind_Identifier)
    if !ok {
        return
    }
    node.Name     = nameTk.Value
    node.Location = nameTk.Location()
    return node, true
}
