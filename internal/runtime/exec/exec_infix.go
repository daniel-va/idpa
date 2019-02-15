package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func execInfix(ctx Context, node ast.InfixNode, mapping InfixMapping) (Value, *Error) {
    lValue, rValue, err := execInfixValues(ctx, node)
    if err != nil {
        return nil, err
    }

    switch lValue.(type) {
    case NumberValue:
        if mapping.Number == nil {
            break
        }
        lNumberValue     := lValue.(NumberValue)
        rNumberValue, ok := rValue.(NumberValue)
        if !ok {
            break
        }
        return mapping.Number(float64(lNumberValue), float64(rNumberValue))
    case StringValue:
        if mapping.String == nil {
            break
        }
        lStringValue     := lValue.(StringValue)
        rStringValue, ok := rValue.(StringValue)
        if !ok {
            break
        }
        return mapping.String(string(lStringValue), string(rStringValue))
    }

    return nil, Err(mapping.Message, lValue.Type(), rValue.Type()).AtNode(node)
}

type InfixMapping struct {
    Message string

    Number       func(a, b float64) (Value, *Error)
    String       func(a, b string) (Value, *Error)
}

func execInfixValues(ctx Context, node ast.InfixNode) (Value, Value, *Error)  {
    lValue, err := exec(ctx, node.Left())
    if err != nil {
        return nil, nil, err
    }

    rValue, err := exec(ctx, node.Right())
    if err != nil {
        return nil, nil, err
    }

    return lValue, rValue, nil
}