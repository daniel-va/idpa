package ast

type BooleanNode struct {
    location
    Value bool
}

func (n BooleanNode) Dump() string {
    if n.Value {
        return "true"
    } else {
        return "false"
    }
}
