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

    lineBuffer []string
}

func (r *Reader) Read() <-chan rune {
    return r.charCh
}

func (r *Reader) Err() error {
    return r.err
}

func (r *Reader) GetBufferedLine(index int) string {
    return r.lineBuffer[index]
}

func (r *Reader) reader() {
    scanner := bufio.NewScanner(r.source)
    for scanner.Scan() {
        line := scanner.Text()
        r.lineBuffer = append(r.lineBuffer, line)
        for _, char := range []rune(line) {
            r.charCh<- char
        }
        r.charCh<- '\n'
    }
    r.err = scanner.Err()
    close(r.charCh)
}