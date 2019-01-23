package ast

type ClosureNode struct {
    location

    Parameters []NameNode
    Block      BlockNode
}

func (n ClosureNode) Dump() string {
    dump := "("
    for i, arg := range n.Parameters {
        if i > 0 {
            dump += ", "
        }
        dump += arg.Dump()
    }
    dump += ")"
    return dump + " " + n.Block.Dump()
}
