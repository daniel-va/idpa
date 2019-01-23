package ast

import "github.com/daniel-va/idpa/internal/source"

type AndNode struct {
    LeftCondition  Node
    RightCondition Node
}

func (n AndNode) Dump() string {
    return n.LeftCondition.Dump() + " && " + n.RightCondition.Dump()
}

func (n AndNode) Loc() source.Location {
    return source.Location{
        Start: n.LeftCondition.Loc().Start,
        End:   n.RightCondition.Loc().End,
    }
}
