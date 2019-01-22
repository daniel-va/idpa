package source

import (
    "bufio"
    "io"
)

func Read(source io.Reader) *Reader {
    r := &Reader{
        source: source,
        charCh: make(chan rune),
    }
    go r.reader()
    return r
}

type Reader struct {
    source io.Reader
    charCh chan rune
    err    error
}

func (r *Reader) Read() <-chan rune {
    return r.charCh
}

func (r *Reader) Err() error {
    return r.err
}

func (r *Reader) reader() {
    scanner := bufio.NewScanner(r.source)
    for scanner.Scan() {
        line := []rune(scanner.Text())
        for _, char := range line {
            r.charCh<- char
        }
        r.charCh<- '\n'
    }
    r.err = scanner.Err()
    close(r.charCh)
}