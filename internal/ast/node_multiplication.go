package ast

type MultiplicationNode struct {
    infix
}

func NewMultiplicationNode() MultiplicationNode {
    return MultiplicationNode{ infix: newInfix() }
}

func (n MultiplicationNode) Dump() string {
    return dumpInfix(n, "*")
}

func (n MultiplicationNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_MultiplicationAndDivision
}