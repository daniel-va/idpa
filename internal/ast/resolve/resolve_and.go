package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveAnd(ctx Context, leftConditionNode ast.Node) (node ast.AndNode, success bool) {
    node.LeftCondition = leftConditionNode
    if _, ok := ctx.Expect(token.Kind_Operator_And); !ok {
        return
    }

    rightConditionNode, ok := ResolveRoot(ctx)
    if !ok {
        return
    }
    node.RightCondition = rightConditionNode
    return node, true
}
