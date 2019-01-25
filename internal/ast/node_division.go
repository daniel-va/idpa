package ast

type DivisionNode struct {
    infix
}

func NewDivisionNode() DivisionNode {
    return DivisionNode{ infix: newInfix() }
}

func (n DivisionNode) Dump() string {
    return dumpInfix(n, "/")
}

func (n DivisionNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_MultiplicationAndDivision
}