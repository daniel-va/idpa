package token

import "github.com/daniel-va/idpa/internal/source"

func NewLexer(input *source.Reader) *Lexer {
    l := &Lexer{
        outputCh: make(chan Token),
    }
    go l.scanner(input)
    return l
}

type Lexer struct {
    outputCh chan Token
}

func (l *Lexer) Scan() <-chan Token {
    return l.outputCh
}

func (l *Lexer) scanner(input *source.Reader) {
    p := newParser(input.Read())
    for {
        tk, ok := p.Parse()
        if !ok {
            close(l.outputCh)
            return
        }
        if tk.Kind != Kind_Whitespace {
            l.outputCh<- tk
        }
    }
}

