package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveBlock(ctx Context) (node ast.BlockNode, success bool) {
    if openTk, ok := ctx.Expect(token.Kind_Brackets_Curly_Open); ok {
        node.Location.Start = openTk.Pos
    } else {
        return
    }

    var previousNode *ast.Node
    for !ctx.Done() {
        if nextTk, ok := ctx.Peek(); ok && nextTk.Kind == token.Kind_Brackets_Curly_Close {
            break
        }

        innerNode, ok := ResolveRoot(ctx)
        if ok {
            if previousNode != nil {
                ctx.checkFollowingNodes(*previousNode, innerNode)
            }
            previousNode = &innerNode
            node.Nodes   = append(node.Nodes, innerNode)
        }
    }

    if closeTk, ok := ctx.Expect(token.Kind_Brackets_Curly_Close); ok {
        node.Location.End = closeTk.EndPos()
    } else {
        return
    }
    return node, true
}
