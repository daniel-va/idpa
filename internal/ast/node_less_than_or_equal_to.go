package ast

type LessThanOrEqualToNode struct {
    infix
}

func NewLessThanOrEqualToNode() LessThanOrEqualToNode {
    return LessThanOrEqualToNode{ infix: newInfix() }
}

func (n LessThanOrEqualToNode) Dump() string {
    return dumpInfix(n, "<=")
}

func (n LessThanOrEqualToNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Comparison
}