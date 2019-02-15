package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
    "github.com/daniel-va/idpa/internal/runtime/library"
)

type Interpreter struct {
    resolver ast.Resolver
}

func NewInterpreter(resolver ast.Resolver) *Interpreter {
    return &Interpreter{
        resolver: resolver,
    }
}

func (i Interpreter) Run() (exitCode int, errors []Error) {
    var astNodes []ast.Node
    LOOP:
    for {
        select{
        case node := <-i.resolver.Read():
            astNodes = append(astNodes, node)

        case err := <-i.resolver.Err():
            errors = append(errors, Error{
                Message:  err.Message,
                Location: err.Location,
            })
        case <-i.resolver.Done():
            break LOOP
        }
    }
    if len(errors) > 0 {
        exitCode = -1
        return
    }


    scope        := NewScope()
    scope.Members = library.Entries
    _, err := ExecBlock(Context{ ContextScope: scope.Subscope() }, ast.BlockNode{
        Nodes: astNodes,
    })

    if err != nil {
        exitCode = 1
        errors = append(errors, *err)
    }
    return
}