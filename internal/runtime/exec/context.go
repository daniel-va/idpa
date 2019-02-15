package exec

type Context struct {
    *ContextScope
}

func (ctx Context) Sub() Context {
    subscope := ctx.Subscope()
    ctx.ContextScope = subscope
    return ctx
}