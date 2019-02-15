package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecAssignment(ctx Context, node ast.AssignmentNode) *Error {
    resultValue, err := exec(ctx, node.Value)
    if err != nil {
        return err
    }
    ctx.Assign(node.Variable.Name, resultValue)
    return nil
}
