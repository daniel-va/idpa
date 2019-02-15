package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecIf(ctx Context, node ast.IfNode) (Value, *Error) {
    value, _, err := execIf_WithStatus(ctx, node)
    return value, err
}

func ExecElse(ctx Context, node ast.ElseNode) (Value, *Error) {
    value, ok, err := execIf_WithStatus(ctx, node.If)
    if ok || err != nil {
        return value, err
    }
    return ExecBlock(ctx, node.Block)
}

func execIf_WithStatus(ctx Context, node ast.IfNode) (Value, bool, *Error) {
    conditionValue, err := exec(ctx, node.Condition)
    if err != nil {
        return nil, false, err
    }
    booleanValue, ok := conditionValue.(BooleanValue)
    if !ok {
        return nil, false, Err("not a Boolean")
    }
    if booleanValue {
        resultValue, err := ExecBlock(ctx, node.Block)
        return resultValue, true, err
    } else {
        return NullValue{}, false, err
    }
}
