package resolve

import "github.com/daniel-va/idpa/internal/ast"

func ResolveNameReference(ctx Context) (node ast.NameReferenceNode, success bool) {
    nameNode, ok := ResolveName(ctx)
    if !ok {
        return
    }
    node.NameNode = nameNode
    return node, true
}