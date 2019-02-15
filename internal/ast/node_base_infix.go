package ast

import (
    "github.com/daniel-va/idpa/internal/source"
)

type InfixNode interface {
    Node

    Precedence() OperatorPrecedence

    Left() Node
    LeftP() *Node

    Right() Node
    RightP() *Node
}

type OperatorPrecedence int8
const (
    _ OperatorPrecedence = iota

    OperatorPrecedence_Or
    OperatorPrecedence_And
    OperatorPrecedence_Equality
    OperatorPrecedence_Comparison
    OperatorPrecedence_AdditionAndSubtraction
    OperatorPrecedence_MultiplicationAndDivision
    OperatorPrecedence_Not
    OperatorPrecedence_NegativeAndPositive
)

type infix struct {
    LeftNode  *Node
    RightNode *Node
}

func newInfix() infix {
    var leftNode Node
    var rightNode Node
    return infix{
        LeftNode: &leftNode,
        RightNode: &rightNode,
    }
}

func (n infix) Left() Node {
    return *n.LeftP()
}

func (n infix) LeftP() *Node {
    return n.LeftNode
}

func (n infix) Right() Node {
    return *n.RightP()
}

func (n infix) RightP() *Node {
    return n.RightNode
}

func (n infix) Loc() source.Location {
    return source.Location{
        Start: n.Left().Loc().Start,
        End:   n.Right().Loc().End,
    }
}

func dumpInfix(node InfixNode, operator string) string {
    return "(" + node.Left().Dump() + " " + operator + " " + node.Right().Dump() + ")"
}