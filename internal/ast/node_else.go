package ast

import "github.com/daniel-va/idpa/internal/source"

type ElseNode struct {
    If    IfNode
    Block BlockNode
}

func (n ElseNode) Dump() string {
    return n.If.Dump() + " else " + n.Block.Dump()
}

func (n ElseNode) Loc() source.Location {
    loc    := n.If.Location
    loc.End = n.Block.Location.End
    return loc
}