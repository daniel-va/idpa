package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
    "math/big"
)


func ResolveInteger(ctx Context) (node ast.IntegerNode, success bool) {
    var isNegative bool
    if prefixTk, ok := ctx.Peek(); ok {
        switch prefixTk.Kind {
        case token.Kind_Operator_Add:
            ctx.Drop()
        case token.Kind_Operator_Subtract:
            ctx.Drop()
            isNegative = true
        }
    }

    valueTk, ok := ctx.Expect(token.Kind_Number)
    if !ok {
        return
    }

    value, ok := (&big.Int{}).SetString(valueTk.Value, 10)
    if !ok {
        ctx.Report(Err("not a valid number").AtToken(valueTk))
        return
    }

    if isNegative {
        value = value.Neg(value)
    }

    if !value.IsInt64() {
        ctx.Report(Err("number too large").AtToken(valueTk))
    }

    node.Location = valueTk.Location()
    node.Value    = value.Int64()
    return node, true
}
