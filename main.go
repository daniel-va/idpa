package main

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/source"
    "github.com/daniel-va/idpa/internal/token"
    "os"
)

func main() {
    if err := Run(); err != nil {
        panic(err)
    }
}

func Run() error {
    file, err := os.Open("./code.txt")
    if err != nil {
        return fmt.Errorf("failed to open source file: %s", err)
    }
    defer file.Close()

    reader := source.Read(file)
    lexer  := token.NewLexer(reader)

    for entry := range lexer.Scan() {
        fmt.Println(entry)
    }

    return nil
}