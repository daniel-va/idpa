package resolve

import (
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func Run(lexer *token.Lexer) (<-chan ast.Node, <-chan ast.Error, <-chan struct{}) {
   ctx := newContext(lexer)

   nodeCh := make(chan ast.Node)
   doneCh := make(chan struct{})
   go func() {
       var previousNode *ast.Node
       for !ctx.Done() {
           if node, ok := ResolveRoot(ctx); ok {
               if previousNode != nil {
                   ctx.checkFollowingNodes(*previousNode, node)
               }
               previousNode = &node
               nodeCh<- node
           } else {
               previousNode = nil
           }
       }
       doneCh<- struct{}{}
       close(nodeCh)
       ctx.Close()
   }()
   return nodeCh, ctx.errCh, doneCh
}