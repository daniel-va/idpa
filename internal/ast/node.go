package ast

import "github.com/daniel-va/idpa/internal/source"

type Node interface {
    Loc() source.Location
    Dump() string
}