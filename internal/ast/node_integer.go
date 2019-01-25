package ast

import "strconv"

type IntegerNode struct {
    location
    Value int64
}

func (n IntegerNode) Dump() string {
    return strconv.FormatInt(n.Value, 10)
}