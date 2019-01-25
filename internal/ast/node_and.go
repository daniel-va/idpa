package ast

type AndNode struct {
    infix
}

func NewAndNode() AndNode {
    return AndNode{ infix: newInfix() }
}

func (n AndNode) Dump() string {
    return dumpInfix(n, "&&")
}

func (n AndNode) Precedence() OperatorPrecedence {
    return OperatorPrecedence_And
}