package exec

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecEquals(ctx Context, node ast.InfixNode) (BooleanValue, *Error) {
    lValue, rValue, err := execInfixValues(ctx, node)
    if err != nil {
        return false, err
    }
    return BooleanValue(execEqualsValues(lValue, rValue)), nil
}

func execEqualsValues(lValue, rValue Value) bool {
    if lValue.Type() != rValue.Type() {
        return false
    }
    switch lValue.(type) {
    case ArrayValue:
        lArrayValue := lValue.(ArrayValue)
        rArrayValue := rValue.(ArrayValue)
        if len(lArrayValue) != len(rArrayValue) {
            return false
        }
        for i, lElementValue := range lArrayValue {
            if !execEqualsValues(lElementValue, rArrayValue[i]) {
                return false
            }
        }
        return true
    case BooleanValue:
        lBoolValue := lValue.(BooleanValue)
        rBoolValue := rValue.(BooleanValue)
        return lBoolValue == rBoolValue
    case ClosureValue:
        return false
    default:
        panic(fmt.Sprintf("unchecked type: %T", lValue))
    }
}