package ast

type OrNode struct {
    location

    leftCondition  Node
    rightCondition Node
}

func (n OrNode) Dump() string {
    return n.leftCondition.Dump() + " || " + n.rightCondition.Dump()
}
