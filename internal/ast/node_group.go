package ast

type GroupNode struct {
    location
    Node Node
}

func (n GroupNode) Dump() string {
    return "(" + n.Node.Dump() + ")"
}