package ast

type Resolver interface {
    Read() <-chan Node
    Err()  <-chan Error
    Done() <-chan struct{}
}
