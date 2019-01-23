package ast

type WhileNode struct {
    location

    Condition Node
    Block     BlockNode
}

func (n WhileNode) Dump() string {
    return "while " + n.Condition.Dump() + " " + n.Block.Dump()
}
