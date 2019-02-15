package ast

type ReturnNode struct {
    location

    Value *Node
}

func (r ReturnNode) Dump() string {
    dump := "return"
    if r.Value != nil {
        dump += " " + (*r.Value).Dump()
    }
    return dump
}