package resolve

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func newContext(lexer *token.Lexer) Context {
    return Context{
        contextInput: &contextInput{
          lexer: lexer,
        },
        errCh: make(chan ast.Error),
    }
}

type Context struct {
    *contextInput
    errCh chan ast.Error
}

func (ctx Context) Expect(kind token.Kind) (token.Token, bool) {
    tk, ok := ctx.Read()
    if !ok {
        ctx.Report(Err("missing %s", kind).AtToken(tk))
        return tk, false
    }
    if tk.Kind != kind {
        ctx.Report(Err("expected %s, got `%s`", kind, tk.Value).AtToken(tk))
        return tk, false
    }
    return tk, true
}

func (ctx Context) Report(err ast.Error) {
    ctx.errCh<- err
}

func (ctx Context) Close() error {
    close(ctx.errCh)
    return nil
}

func (ctx Context) checkFollowingNodes(node1, node2 ast.Node) bool {
    if node1.Loc().End.Row == node2.Loc().Start.Row {
        ctx.Report(
            Err("consecutive statements must be on separate lines").
                AtNode(node2),
        )
        return false
    }
    return true
}

func Err(format string, elements ...interface{}) ast.Error {
    return ast.Error{
        Message: fmt.Sprintf(format, elements...),
    }
}