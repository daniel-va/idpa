package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveAnd(ctx Context, leftConditionNode ast.Node) (resultNode ast.Node, success bool) {
    node := ast.AndNode{}
    node.LeftCondition = leftConditionNode
    if _, ok := ctx.Expect(token.Kind_Operator_And); !ok {
        return
    }

    rightConditionNode, ok := ResolveRoot(ctx)
    if !ok {
        return
    }

    if orNode, ok := rightConditionNode.(ast.OrNode); ok {
        // operator precedence:
        // && > ||

        node.RightCondition  = orNode.LeftCondition
        orNode.LeftCondition = node
        return orNode, true
    } else {
        node.RightCondition = rightConditionNode
        return node, true
    }
}
