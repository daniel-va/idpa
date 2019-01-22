package ast

type NameNode struct {
    location
    Name string
}

func (n NameNode) Dump() string {
    return n.Name
}


