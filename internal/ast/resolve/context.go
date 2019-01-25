package resolve

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/token"
)

func newContext(lexer *token.Lexer, resolver ast.ProxyResolver) Context {
    ctx := Context{
        contextInput: &contextInput{
          lexer: lexer,
        },
        resolver: resolver,
    }
    ctx.Peek() // check if already done (in case of empty file)
    return ctx
}

type Context struct {
    *contextInput
    resolver ast.ProxyResolver
}

func (ctx Context) ExpectAny() (token.Token, bool) {
    tk, ok := ctx.Read()
    if !ok {
        ctx.Report(Err("missing token").AtToken(tk))
    }
    return tk, ok
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
    ctx.resolver.SendErr(err)
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