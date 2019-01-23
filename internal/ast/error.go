package ast

import (
    "github.com/daniel-va/idpa/internal/source"
    "github.com/daniel-va/idpa/internal/token"
)

type Error struct {
    Message  string
    Location source.Location
}

func (e Error) AtToken(tk token.Token) Error {
    e.Location = tk.Location()
    return e
}

func (e Error) AtNode(node Node) Error {
    e.Location = node.Loc()
    return e
}

func (e Error) From(pos source.Pos) Error {
    e.Location.Start = pos
    return e
}

func (e Error) To(pos source.Pos) Error {
    e.Location.End = pos
    return e
}