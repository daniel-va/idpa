package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveString(ctx Context) (node ast.StringNode, ok bool) {
    stringTk, ok := ctx.Expect(token.Kind_String)
    if !ok {
        return
    }
    stringLen := len(stringTk.Value)
    if stringLen < 2 || stringTk.Value[stringLen - 1] != '"' {
        ctx.Report(Err("unclosed string literal").AtToken(stringTk))
    }
    node.Value    = stringTk.Value
    node.Location = stringTk.Location()
    return node, true
}
