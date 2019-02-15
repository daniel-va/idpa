package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecClosure(ctx Context, node ast.ClosureNode) (Value, *Error) {
    return ClosureValue{
        Arity: len(node.Parameters),
        Call:  func(params []Value) (Value, *Error) {
            subCtx := ctx.Sub()
            for i, param := range node.Parameters {
                subCtx.Assign(param.Name, params[i])
            }
            return ExecBlock(subCtx, node.Block)
        },
    }, nil
}
