package ast

import (
    "strconv"
)

type FloatNode struct {
    location
    Value float64
}

func (n FloatNode) Dump() string {
    return strconv.FormatFloat(n.Value, 'f', -1, 64)
}