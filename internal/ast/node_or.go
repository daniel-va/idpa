package ast

type OrNode struct {
    location

    LeftCondition  Node
    RightCondition Node
}

func (n OrNode) Dump() string {
    return "(" + n.LeftCondition.Dump() + " || " + n.RightCondition.Dump() + ")"
}
