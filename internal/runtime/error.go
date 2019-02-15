package runtime

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/source"
)

type Error struct {
    Message  string
    Location source.Location
}

func Err(format string, elements ...interface{}) *Error {
    return &Error{
        Message: fmt.Sprintf(format, elements...),
    }
}

func (e *Error) AtNode(node ast.Node) *Error {
    e.Location = node.Loc()
    return e
}