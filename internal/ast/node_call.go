package ast

type CallNode struct {
    location

    Target    Node
    Arguments []Node
}

func (n CallNode) Dump() string {
    dump := n.Target.Dump() + "("
    for i, arg := range n.Arguments {
        if i > 0 {
            dump += ", "
        }
        dump += arg.Dump()
    }
    return dump + ")"
}
