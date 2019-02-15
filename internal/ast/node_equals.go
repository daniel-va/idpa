package ast

type EqualsNode struct {
    infix
}

func NewEqualsNode() EqualsNode {
    return EqualsNode{ infix: newInfix() }
}

func (n EqualsNode) Dump() string {
    return dumpInfix(n, "==")
}

func (n EqualsNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Equality
}
