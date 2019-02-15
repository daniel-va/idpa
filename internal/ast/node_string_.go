package ast

import "strconv"

type StringNode struct {
    location
    Value string
}

func (n StringNode) Dump() string {
    return strconv.Quote(n.UsableValue())
}

func (n StringNode) UsableValue() string {
    return n.Value[1:len(n.Value) - 1]
}
