package source

import "fmt"

type Pos struct {
    Row, Col int
}

func (p Pos) AddCol(amount int) Pos {
    p.Col += amount
    return p
}

func (p Pos) AddRow(amount int) Pos {
    p.Row += amount
    return p
}

func (p Pos) String() string {
    return fmt.Sprint(p.Row) + ":" + fmt.Sprint(p.Col)
}

func (p Pos) AtPath(path string) string {
    return path + ":" + p.String()
}