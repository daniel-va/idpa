package ast

import "strings"

type BlockNode struct{
    location
    Nodes []Node
}

func (n BlockNode) Dump() string {
    var sep string
    if len(n.Nodes) > 0 {
        sep = "\n"
    }

    dump := ""
    for i, node := range n.Nodes {
        if i > 0 {
            dump += "\n"
        }
        dump += "\t" + strings.Replace(node.Dump(), "\n", "\n\t", -1)
    }
    return "{" + sep + dump + sep + "}"
}