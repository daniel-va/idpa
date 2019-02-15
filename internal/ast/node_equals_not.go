package ast

type NotEqualsNode struct {
    infix
}

func NewNotEqualsNode() NotEqualsNode {
    return NotEqualsNode{ infix: newInfix() }
}

func (n NotEqualsNode) Dump() string {
    return dumpInfix(n, "!=")
}

func (n NotEqualsNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Equality
}
