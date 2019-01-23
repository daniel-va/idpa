package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveElse(ctx Context, ifNode ast.IfNode) (node ast.ElseNode, success bool) {
    if _, ok := ctx.Expect(token.Kind_Keyword_Else); !ok {
        return
    }

    blockNode, ok := ResolveBlock(ctx)
    if !ok {
        return
    }

    node.If    = ifNode
    node.Block = blockNode
    return node, true
}
