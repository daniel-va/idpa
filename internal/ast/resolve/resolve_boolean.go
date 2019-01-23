package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveBoolean(ctx Context) (node ast.BooleanNode, success bool) {
    boolTk, ok := ctx.Read()
    if !ok {
        ctx.Report(Err("not a Boolean literal").AtToken(boolTk))
        return
    }
    node.Location = boolTk.Location()
    switch boolTk.Kind {
    case token.Kind_Keyword_True:
        node.Value = true
    case token.Kind_Keyword_False:
        node.Value = false
    default:
        return
    }
    return node, true
}