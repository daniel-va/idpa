package resolve

import (
    "github.com/daniel-va/idpa/internal/source"
    "github.com/daniel-va/idpa/internal/token"
)

type contextInput struct {
    lexer   *token.Lexer
    lastPos source.Pos
    done    bool

    peeks []token.Token
}

func (i *contextInput) Read() (token.Token, bool) {
    if len(i.peeks) > 0 {
        entry := i.peeks[0]
        i.Drop()
        return entry, true
    }
    return i.next()
}

func (i *contextInput) Peek() (token.Token, bool) {
    if len(i.peeks) > 0 {
        return i.peeks[0], true
    }
    peeks, ok := i.PeekAll(1)
    if !ok {
        return i.Read()
    }
    return peeks[0], true
}

func (i *contextInput) PeekAll(amount int) ([]token.Token, bool) {
    if count := len(i.peeks); count < amount {
        for count < amount {
            entry, ok := i.next()
            if !ok {
                return nil, false
            }
            count += 1
            i.peeks = append(i.peeks, entry)
        }
    }
    return i.peeks[:amount], true
}

func (i *contextInput) Done() bool {
    return i.done
}

func (i *contextInput) Drop() {
    i.peeks = i.peeks[1:]
}

func (i *contextInput) Push(tk token.Token) {
    i.peeks = append([]token.Token{ tk }, i.peeks...)
}

func (i *contextInput) Pos() source.Pos {
    entry, _ := i.Peek()
    return entry.Pos
}

func (i *contextInput) next() (token.Token, bool) {
    entry, ok := <-i.lexer.Scan()
    if !ok {
        i.done = true
        return token.Token{ Pos: i.lastPos.AddCol(1) }, false
    }
    i.lastPos = entry.EndPos()
    return entry, true
}