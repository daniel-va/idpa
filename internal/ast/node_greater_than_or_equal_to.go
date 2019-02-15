package ast

type GreaterThanOrEqualToNode struct {
    infix
}

func NewGreaterThanOrEqualToNode() GreaterThanOrEqualToNode {
    return GreaterThanOrEqualToNode{ infix: newInfix() }
}

func (n GreaterThanOrEqualToNode) Dump() string {
    return dumpInfix(n, ">=")
}

func (n GreaterThanOrEqualToNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Comparison
}