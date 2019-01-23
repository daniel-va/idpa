package ast

type IfNode struct {
    location

    Condition Node
    Block     BlockNode
}

func (n IfNode) Dump() string {
    return "if " + n.Condition.Dump() + " " + n.Block.Dump()
}