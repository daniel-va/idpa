package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecCall(ctx Context, node ast.CallNode) (Value, *Error) {
    targetValue, err := exec(ctx, node.Target)
    if err != nil {
       return nil, err
    }

    closureValue, ok := targetValue.(ClosureValue)
    if !ok {
        return nil, Err("not a callable value").AtNode(node)
    }

    var arity int
    var varargs bool
    if closureValue.Arity == -1 {
        arity   = len(node.Arguments)
        varargs = true
    } else {
        arity = closureValue.Arity
        if arity != len(node.Arguments) {
            return nil, Err("argument mismatch, expected %d, got %d", arity, len(node.Arguments)).AtNode(node)
        }
    }

    paramTypes := closureValue.ParamTypes
    if paramTypes != nil {
        typeLen := len(paramTypes)
        if (varargs && typeLen != -1) || (!varargs && typeLen != arity) {
            panic("ParamTypes do not match closure arity")
        }
    }

    argValues := make([]Value, arity)
    for i, argNode:= range node.Arguments {
        argValue, err := exec(ctx, argNode)
        if err != nil {
            return nil, err
        }
        if paramTypes != nil {
            var paramType TypeAnnotation
            if varargs {
                paramType = paramTypes[0]
            } else {
                paramType = paramTypes[i]
            }
            if paramType != 0 && !paramType.Allows(argValue.Type()) {
                return nil, Err("value of type %s not allowed here", argValue.Type()).AtNode(node)
            }
        }
        argValues[i] = argValue
    }

    resultValue, err := closureValue.Call(argValues)
    if err != nil && err.Location.Start.Row == 0 {
        err = err.AtNode(node)
    }
    return resultValue, err
}