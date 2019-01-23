package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveClosure(ctx Context) (resultNode ast.Node, success bool) {
    node := ast.ClosureNode{}

    if openTk, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Open); ok {
        node.Location.Start = openTk.Pos
    } else {
        return
    }

    canBeGroup := true

    for !ctx.Done() {
        if nextTk, ok := ctx.Peek(); ok && nextTk.Kind == token.Kind_Brackets_Parentheses_Close {
            break
        }

        if !canBeGroup {
            ctx.Expect(token.Kind_Syntax_ValueSeparator)
        }

        if canBeGroup {
            paramNode, ok := ResolveRoot(ctx)
            if !ok {
                break
            }

            node.Parameters = append(node.Parameters, paramNode)
        }


    }

    if closeTk, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Close); ok {
        node.Location.End = closeTk.EndPos()
    }

    if nextTk, ok := ctx.Peek(); ok && nextTk.Kind == token.Kind_Brackets_Curly_Open {
        closureNode := ast.ClosureNode{
            Parameters: parameters,
        }
        closureNode.Location.Start = node.Location.Start

        blockNode, ok := ResolveBlock(ctx)
        if !ok {
            return
        }
        closureNode.Block        = blockNode
        closureNode.Location.End = blockNode.Location.End
        return closureNode, true
    }
}
