package ast

import (
    "github.com/daniel-va/idpa/internal/source"
)

type AssignmentNode struct {
    Variable NameNode
    Value    Node
}

func (n AssignmentNode) Dump() string {
    return n.Variable.Dump() + " = " + n.Value.Dump()
}

func (n AssignmentNode) Loc() source.Location {
    return source.Location{
        Start: n.Variable.Loc().Start,
        End:   n.Value.Loc().End,
    }
}

