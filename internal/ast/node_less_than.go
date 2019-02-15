package ast

type LessThanNode struct {
    infix
}

func NewLessThanNode() LessThanNode {
    return LessThanNode{ infix: newInfix() }
}

func (n LessThanNode) Dump() string {
    return dumpInfix(n, "<")
}

func (n LessThanNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Comparison
}