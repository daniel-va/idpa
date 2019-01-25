package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
)

func ResolveInfix(ctx Context, leftNode ast.Node, node ast.InfixNode) (resultNode ast.InfixNode, success bool) {
    *node.LeftP() = leftNode
    if _, ok := ctx.ExpectAny(); !ok {
        return
    }

    rightNode, ok := ResolveRoot(ctx)
    if !ok {
        return
    }
    *node.RightP() = rightNode
    return orderInfix(node), true
}
