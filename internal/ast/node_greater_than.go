package ast

type GreaterThanNode struct {
    infix
}

func NewGreaterThanNode() GreaterThanNode {
    return GreaterThanNode{ infix: newInfix() }
}

func (n GreaterThanNode) Dump() string {
    return dumpInfix(n, ">")
}

func (n GreaterThanNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Comparison
}