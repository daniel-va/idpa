package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveAssignment(ctx Context, variableNode ast.NameReferenceNode) (node ast.AssignmentNode, success bool) {
    if _, ok := ctx.Expect(token.Kind_Operator_Assign); !ok {
        return
    }

    valueNode, ok := ResolveRoot(ctx)
    if !ok {
        return
    }
    node.Variable = variableNode
    node.Value    = valueNode
    return node, true
}
