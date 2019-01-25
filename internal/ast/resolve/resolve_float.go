package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
    "math/big"
)

func ResolveFloat(ctx Context, integerNode ast.IntegerNode) (node ast.FloatNode, success bool) {
    if _, ok := ctx.Expect(token.Kind_Syntax_DecimalPoint); !ok {
        return
    }

    decimalTk, ok := ctx.Expect(token.Kind_Number)
    if !ok {
        return
    }

    node.Location.Start = integerNode.Loc().Start
    node.Location.End   = decimalTk.EndPos()

    bigValue, ok := big.NewFloat(0).SetString(integerNode.Dump() + "." + decimalTk.Value)
    if !ok {
        ctx.Report(Err("not a valid number").AtNode(node))
        return
    }

    value, accuracy := bigValue.Float64()
    if accuracy != big.Exact {
        ctx.Report(Err("number too large").AtNode(node))
    }

    node.Value = value
    return node, true
}
