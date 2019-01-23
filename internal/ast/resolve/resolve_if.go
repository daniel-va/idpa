package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveIf(ctx Context) (node ast.IfNode, success bool) {
    if ifTk, ok := ctx.Expect(token.Kind_Keyword_If); ok {
        node.Location.Start = ifTk.Pos
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
