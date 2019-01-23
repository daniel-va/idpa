package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveCall(ctx Context, targetNode ast.Node) (node ast.CallNode, success bool) {
    node.Target         = targetNode
    node.Location.Start = targetNode.Loc().Start
    if _, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Open); !ok {
        return
    }

    for !ctx.Done() {
        if nextTk, ok := ctx.Peek(); ok && nextTk.Kind == token.Kind_Brackets_Parentheses_Close {
            break
        }

        if len(node.Arguments) > 0 {
            ctx.Expect(token.Kind_Syntax_ValueSeparator)
        }

        argNode, ok := ResolveRoot(ctx)
        if ok {
            node.Arguments = append(node.Arguments, argNode)
        }
    }

    if closeTk, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Close); ok {
        node.Location.End = closeTk.EndPos()
    } else {
        return
    }

    return node, true
}
