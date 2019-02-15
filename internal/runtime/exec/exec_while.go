package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecWhile(ctx Context, node ast.WhileNode) (Value, *Error) {
    execCondition := func() (bool, *Error) {
        conditionValue, err := exec(ctx, node.Condition)
        if err != nil {
            return false, err
        }

        boolValue, ok := conditionValue.(BooleanValue)
        if !ok {
            return false, Err("condition must be a Boolean").AtNode(node.Condition)
        }

        return bool(boolValue), nil
    }

    var result ArrayValue
    for {
        ok, err := execCondition()
        if err != nil {
            return nil, err
        }
        if !ok {
            break
        }

        resultValue, _, err := execAny(ctx, node.Block)
        if err != nil {
            return nil, err
        }
        result = append(result, resultValue)
    }

    return result, nil
}