package token

import "github.com/daniel-va/idpa/internal/source"

type Token struct {
    Value string
    Kind  Kind
    Pos   source.Pos
}

func (t Token) EndPos() source.Pos {
    return t.Pos.AddCol(len(t.Value))
}

func (t Token) Location() source.Location {
    return source.Location{
        Start: t.Pos,
        End:   t.EndPos(),
    }
}

func (t Token) String() string {
    return t.Kind.String() + "(" + t.Value + ")[" + t.Pos.String() + "]"
}