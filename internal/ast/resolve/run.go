package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func Run(lexer *token.Lexer) ast.Resolver {
    resolver := ast.NewResolver()
    ctx := newContext(lexer, resolver)
    go func() {
       var previousNode *ast.Node
       for !ctx.Done() {
           if node, ok := ResolveRoot(ctx); ok {
               if previousNode != nil {
                   ctx.checkFollowingNodes(*previousNode, node)
               }
               previousNode = &node
               resolver.Send(node)
           } else {
               previousNode = nil
           }
       }
       resolver.Close()
    }()
    return resolver
}