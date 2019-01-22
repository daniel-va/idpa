package ast

import "github.com/daniel-va/idpa/internal/source"

type location struct {
    Location source.Location
}

func (n location) Loc() source.Location {
    return n.Location
}