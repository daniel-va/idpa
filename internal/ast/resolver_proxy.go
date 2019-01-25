package ast

type ProxyResolver struct {
    nodeCh chan Node
    errCh  chan Error
    doneCh chan struct{}
}

func NewResolver() ProxyResolver {
    return ProxyResolver{
        nodeCh: make(chan Node),
        errCh:  make(chan Error),
        doneCh: make(chan struct{}),
    }
}

func (r ProxyResolver) Send(node Node) {
    r.nodeCh<- node
}

func (r ProxyResolver) SendErr(err Error) {
    r.errCh<- err
}

func (r ProxyResolver) Read() <-chan Node {
    return r.nodeCh
}

func (r ProxyResolver) Err() <-chan Error {
    return r.errCh
}

func (r ProxyResolver) Done() <-chan struct{} {
    return r.doneCh
}

func (r ProxyResolver) Close() error {
    r.doneCh<- struct{}{}
    close(r.doneCh)
    close(r.errCh)
    close(r.nodeCh)
    return nil
}