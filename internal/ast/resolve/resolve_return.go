package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveReturn(ctx Context) (node ast.ReturnNode, success bool) {
    if returnTk, ok := ctx.Expect(token.Kind_Keyword_Return); ok {
        node.Location = returnTk.Location()
    } else {
        return
    }

    ctx.Report(Err("return instructions are not supported yet."))
    return node, true
}