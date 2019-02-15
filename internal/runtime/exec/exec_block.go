package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecBlock(ctx Context, node ast.BlockNode) (Value, *Error) {
    var resultValue Value = NullValue{}
    for _, lineNode := range node.Nodes {
        lineValue, isExpression, err := execAny(ctx, lineNode)
        if err != nil {
            return nil, err
        }
        if isExpression {
            resultValue = lineValue
        } else {
            resultValue = NullValue{}
        }
    }
    return resultValue, nil
}
