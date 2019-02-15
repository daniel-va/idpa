package exec

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func exec(ctx Context, node ast.Node) (Value, *Error) {
    value, isExpression, err := execAny(ctx, node)
    if err != nil {
        return nil, err
    }
    if !isExpression {
        return nil, Err("statements not allowed here").AtNode(node)
    }
    return value, nil
}

func execStatement(ctx Context, node ast.Node) *Error {
    _, _, err := execAny(ctx, node)
    return err
}

func execAny(ctx Context, node ast.Node) (Value, bool, *Error) {
    expr := func(value Value, err *Error) (Value, bool, *Error) {
        return value, true, err
    }

    stmt := func(err *Error) (Value, bool, *Error) {
        return NullValue{}, false, err
    }

    switch node.(type) {
    case ast.AdditionNode:
        return expr(ExecAddition(ctx, node.(ast.AdditionNode)))
    case ast.AndNode:
        return expr(ExecAnd(ctx, node.(ast.AndNode)))
    case ast.AssignmentNode:
        return stmt(ExecAssignment(ctx, node.(ast.AssignmentNode)))
    case ast.BlockNode:
        return expr(ExecBlock(ctx, node.(ast.BlockNode)))
    case ast.BooleanNode:
        return expr(ExecBoolean(ctx, node.(ast.BooleanNode)))
    case ast.ClosureNode:
        return expr(ExecClosure(ctx, node.(ast.ClosureNode)))
    case ast.CallNode:
        return expr(ExecCall(ctx, node.(ast.CallNode)))
    case ast.DivisionNode:
        return expr(ExecDivision(ctx, node.(ast.DivisionNode)))
    case ast.ElseNode:
        return expr(ExecElse(ctx, node.(ast.ElseNode)))
    case ast.EqualsNode:
        return expr(ExecEquals(ctx, node.(ast.EqualsNode)))
    case ast.NotEqualsNode:
        return expr(ExecNotEquals(ctx, node.(ast.NotEqualsNode)))
    case ast.FloatNode:
        return expr(ExecFloat(ctx, node.(ast.FloatNode)))
    case ast.GreaterThanNode:
        return expr(ExecGreaterThan(ctx, node.(ast.GreaterThanNode)))
    case ast.GreaterThanOrEqualToNode:
        return expr(ExecGreaterThanOrEqualTo(ctx, node.(ast.GreaterThanOrEqualToNode)))
    case ast.GroupNode:
        return expr(ExecGroup(ctx, node.(ast.GroupNode)))
    case ast.IfNode:
        return expr(ExecIf(ctx, node.(ast.IfNode)))
    case ast.IntegerNode:
        return expr(ExecInteger(ctx, node.(ast.IntegerNode)))
    case ast.LessThanNode:
        return expr(ExecLessThan(ctx, node.(ast.LessThanNode)))
    case ast.LessThanOrEqualToNode:
        return expr(ExecLessThanOrEqualTo(ctx, node.(ast.LessThanOrEqualToNode)))
    case ast.MultiplicationNode:
        return expr(ExecMultiplication(ctx, node.(ast.MultiplicationNode)))
    case ast.NameNode:
        return expr(ExecName(ctx, node.(ast.NameNode)))
    case ast.NameReferenceNode:
        return expr(ExecNameReference(ctx, node.(ast.NameReferenceNode)))
    case ast.OrNode:
        return expr(ExecOr(ctx, node.(ast.OrNode)))
    case ast.StringNode:
        return expr(ExecString(ctx, node.(ast.StringNode)))
    case ast.SubtractionNode:
        return expr(ExecSubtraction(ctx, node.(ast.SubtractionNode)))
    case ast.WhileNode:
        return expr(ExecWhile(ctx, node.(ast.WhileNode)))
    default:
        panic(fmt.Sprintf("unchecked type: %T", node))
    }
}