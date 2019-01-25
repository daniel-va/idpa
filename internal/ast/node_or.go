package ast

type OrNode struct {
    infix
}

func NewOrNode() OrNode {
    return OrNode{ infix: newInfix() }
}

func (n OrNode) Dump() string {
    return dumpInfix(n, "||")
}

func (n OrNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_Or
}