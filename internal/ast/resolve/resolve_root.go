package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func ResolveRoot(ctx Context) (ast.Node, bool) {
    node, ok := root_base(ctx)
    if !ok {
        return node, ok
    }

    for {
        nextNode, ok := root_followUp(ctx, node)
        if !ok {
            return node, true
        }
        node = nextNode
    }
}

func root_base(ctx Context) (ast.Node, bool) {
    tk, ok := ctx.Peek()
    if !ok {
        // should always be an expression that's missing,
        // since statements are at the top-level.
        // They are already filtered out by the `!ctx.Done()` in `run()`.
        ctx.Report(Err("missing expression").AtToken(tk))
        return nil, false
    }
    switch tk.Kind {
    case token.Kind_Identifier:
        return ResolveNameReference(ctx)
    case token.Kind_String:
        return ResolveString(ctx)
    case token.Kind_Keyword_True, token.Kind_Keyword_False:
        return ResolveBoolean(ctx)
    case token.Kind_Brackets_Curly_Open:
        return ResolveBlock(ctx)
    case token.Kind_Brackets_Parentheses_Open:
        return ResolveGroup(ctx)
    case token.Kind_Keyword_If:
        return ResolveIf(ctx)
    case token.Kind_Keyword_While:
        return ResolveWhile(ctx)
    default:
        ctx.Drop()
        ctx.Report(Err("unexpected `%s`", tk.Value).AtToken(tk))
        return ResolveRoot(ctx)
    }
}

func root_followUp(ctx Context, startNode ast.Node) (node ast.Node, success bool) {
    nextTk, ok := ctx.Peek()
    if !ok {
        return
    }
    switch nextTk.Kind {
    case token.Kind_Operator_Assign:
        variableNode, isValidTarget:= startNode.(ast.NameReferenceNode)
        node, ok := ResolveAssignment(ctx, variableNode)
        if !isValidTarget {
            ctx.Report(Err("assignment to non-variable").From(startNode.Loc().Start).To(nextTk.Pos))
        }
        return node, ok
    case token.Kind_Operator_And:
        return ResolveAnd(ctx, startNode)
    case token.Kind_Operator_Or:
        return ResolveOr(ctx, startNode)
    case token.Kind_Brackets_Parentheses_Open:
        return ResolveCall(ctx, startNode)
    case token.Kind_Keyword_Else:
        ifNode, ok := startNode.(ast.IfNode)
        if !ok {
            return
        }
        return ResolveElse(ctx, ifNode)
    default:
        return
    }
}