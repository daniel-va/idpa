package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func execInfixBoolean(
    ctx Context, node ast.InfixNode, action func(a bool, execB func() (bool, *Error)) (bool, *Error),
) (Value, *Error) {
    errNotABoolean := func(value Value) *Error {
        return Err("expected %s, got %s", ValueType_Boolean, value.Type())
    }

    lValue, err := exec(ctx, node.Left())
    if err != nil {
        return nil, err
    }

    lBoolValue, ok := lValue.(BooleanValue)
    if !ok {
        return nil, errNotABoolean(lBoolValue)
    }

    computeRBoolValue := func() (bool, *Error) {
        rBoolValue, ok := lValue.(BooleanValue)
        if !ok {
            return false, errNotABoolean(lBoolValue)
        }
        return bool(rBoolValue), nil
    }

    result, err := action(bool(lBoolValue), computeRBoolValue)
    return BooleanValue(result), err
}