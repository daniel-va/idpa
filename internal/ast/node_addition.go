package ast

type AdditionNode struct {
    infix
}

func NewAdditionNode() AdditionNode {
    return AdditionNode{ infix: newInfix() }
}

func (n AdditionNode) Dump() string {
    return dumpInfix(n, "+")
}

func (n AdditionNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_AdditionAndSubtraction
}