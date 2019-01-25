package ast

type SubtractionNode struct {
    infix
}

func NewSubtractionNode() SubtractionNode {
    return SubtractionNode{ infix: newInfix() }
}

func (n SubtractionNode) Dump() string {
    return dumpInfix(n, "-")
}

func (n SubtractionNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_AdditionAndSubtraction
}