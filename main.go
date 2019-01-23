package main

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast"
    "github.com/daniel-va/idpa/internal/ast/resolve"
    "github.com/daniel-va/idpa/internal/source"
    "github.com/daniel-va/idpa/internal/token"
    "os"
    "time"
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
    nodeCh, errCh, doneCh := resolve.Run(lexer)
    var errs []ast.Error

    LOOP:
    for {
        select{
        case node := <-nodeCh:
            fmt.Println(node.Dump())

        case err := <-errCh:
            errs = append(errs, err)

        case <-doneCh:
            break LOOP
        }
    }

    // wait for stdout
    time.Sleep(time.Millisecond * 20)
    for _, err := range errs {
        os.Stderr.WriteString(fmt.Sprintf("Error at %s: %s\n", err.Location.Start.AtPath("./code.txt"), err.Message))
    }
    return nil
}