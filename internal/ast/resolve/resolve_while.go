package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveWhile(ctx Context) (node ast.WhileNode, success bool) {
    if whileTk, ok := ctx.Expect(token.Kind_Keyword_While); ok {
        node.Location.Start = whileTk.Pos
    } else {
        return
    }

    conditionNode, ok := ResolveRoot(ctx)
    if !ok {
        return
    }
    node.Condition = conditionNode

    blockNode, ok := ResolveBlock(ctx)
    if !ok {
        return
    }
    node.Block        = blockNode
    node.Location.End = blockNode.Location.End

    return node, true
}
