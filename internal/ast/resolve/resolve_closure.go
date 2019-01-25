package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveClosure(ctx Context) (resultNode ast.Node, success bool) {
    var groupNode ast.GroupNode
    var paramNodes []ast.NameNode


    if openTk, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Open); ok {
        groupNode.Location.Start = openTk.Pos
    } else {
        return
    }

    for !ctx.Done() {
        if groupNode.Node != nil && len(paramNodes) == 0 {
            break
        }

        if nextTk, ok := ctx.Peek(); ok && nextTk.Kind == token.Kind_Brackets_Parentheses_Close {
            break
        }

        if len(paramNodes) > 0 {
            ctx.Expect(token.Kind_Syntax_ValueSeparator)
            contentNode, ok := ResolveRoot(ctx)
            if !ok {
                return
            }
            groupNode.Node = contentNode
            if paramNode, ok := contentNode.(ast.NameReferenceNode); ok {
                paramNodes = append(paramNodes, paramNode.NameNode)
            }
        } else {
            nameNode, ok := ResolveName(ctx)
            if !ok {
                return
            }
            paramNodes = append(paramNodes, nameNode)
        }
    }

    if closeTk, ok := ctx.Expect(token.Kind_Brackets_Parentheses_Close); ok {
        groupNode.Location.End = closeTk.EndPos()
    } else {
        return
    }

    nextTk, ok := ctx.Peek()
    var notAClosure = len(paramNodes) == 0 || !ok || nextTk.Kind != token.Kind_Brackets_Curly_Open
    if notAClosure && groupNode.Node != nil {
       return groupNode, true
    }

    blockNode, ok := ResolveBlock(ctx)
    if !ok {
        return
    }

    node := ast.ClosureNode{
        Parameters: paramNodes,
        Block:      blockNode,
    }
    node.Location.Start = groupNode.Location.Start
    node.Location.End   = blockNode.Location.End
    return node, true
}
