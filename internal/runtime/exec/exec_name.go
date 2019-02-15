package exec

import (
    "github.com/daniel-va/idpa/internal/ast"
    . "github.com/daniel-va/idpa/internal/runtime"
)

func ExecName(ctx Context, node ast.NameNode) (Value, *Error) {
    panic(errNameNodeUsage)
}

const errNameNodeUsage =
`Pure values of NameNode can't appear at runtime.
 They have to be wrapped in a NameNodeReference, if they are supposed to be evaluated.
 This is most likely a compiler bug, and should be passed on to the developers.`
