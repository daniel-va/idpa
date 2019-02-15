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
        return ResolveClosure(ctx)
    case token.Kind_Keyword_If:
        return ResolveIf(ctx)
    case token.Kind_Keyword_While:
        return ResolveWhile(ctx)
    case token.Kind_Number,
         token.Kind_Operator_Subtract,
         token.Kind_Operator_Add:
        return ResolveInteger(ctx)
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
        if !isValidTarget {
            ctx.Report(Err("assignment to non-variable").AtNode(startNode))
        }
        node, ok := ResolveAssignment(ctx, variableNode)
        return node, ok
    case token.Kind_Operator_And:
        return ResolveInfix(ctx, startNode, ast.NewAndNode())
    case token.Kind_Operator_Or:
        return ResolveInfix(ctx, startNode, ast.NewOrNode())
    case token.Kind_Operator_Add:
        return ResolveInfix(ctx, startNode, ast.NewAdditionNode())
    case token.Kind_Operator_Subtract:
        return ResolveInfix(ctx, startNode, ast.NewSubtractionNode())
    case token.Kind_Operator_Multiply:
        return ResolveInfix(ctx, startNode, ast.NewMultiplicationNode())
    case token.Kind_Operator_Divide:
        return ResolveInfix(ctx, startNode, ast.NewDivisionNode())
    case token.Kind_Operator_GreaterThan:
        return ResolveInfix(ctx, startNode, ast.NewGreaterThanNode())
    case token.Kind_Operator_GreaterThanOrEqualTo:
        return ResolveInfix(ctx, startNode, ast.NewGreaterThanOrEqualToNode())
    case token.Kind_Operator_LessThan:
        return ResolveInfix(ctx, startNode, ast.NewLessThanNode())
    case token.Kind_Operator_LessThanOrEqualTo:
        return ResolveInfix(ctx, startNode, ast.NewLessThanOrEqualToNode())
    case token.Kind_Operator_Equal:
        return ResolveInfix(ctx, startNode, ast.NewEqualsNode())
    case token.Kind_Operator_NotEqual:
        return ResolveInfix(ctx, startNode, ast.NewNotEqualsNode())
    case token.Kind_Brackets_Parentheses_Open:
        return ResolveCall(ctx, startNode)
    case token.Kind_Keyword_Else:
        ifNode, ok := startNode.(ast.IfNode)
        if !ok {
            return
        }
        return ResolveElse(ctx, ifNode)
    case token.Kind_Syntax_DecimalPoint:
        integerNode, ok := startNode.(ast.IntegerNode)
        if !ok {
            return
        }
        return ResolveFloat(ctx, integerNode)
    default:
        return
    }
}

func orderInfix(node ast.InfixNode) ast.InfixNode {
    precedence := node.Precedence()
    if rightNode, ok := node.Right().(ast.InfixNode); ok {
        otherPrecedence :=  rightNode.Precedence()
        if otherPrecedence <= precedence {
            *node.RightP()     = rightNode.Left()
            *rightNode.LeftP() = orderInfix(node)
            return orderInfix(rightNode)
        }
    }
    return node
}