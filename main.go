package main

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/ast/resolve"
    "github.com/daniel-va/idpa/internal/runtime/exec"
    "github.com/daniel-va/idpa/internal/source"
    "github.com/daniel-va/idpa/internal/token"
    "github.com/urfave/cli"
    "os"
    "strconv"
    "strings"
)

func main() {
    app        := cli.NewApp()
    app.Version = "1.0.0"
    app.Name    = "idpa"
    app.Usage   = "the interpreter for the idpa-linux programming language"

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "file, f",
            Usage: "the `file` to interpret",
        },
    }

    app.Action = func(c *cli.Context) error {
        file := c.String("file")
        if file == "" {
            return fmt.Errorf("missing `--file` flag")
        }
        return Run(file)
    }

    if err := app.Run(os.Args); err != nil {
        os.Stderr.WriteString("[ERROR] " + err.Error() + "\n")
    }
}

const codeColOffset = 4

func Run(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open source file: %s", err)
    }
    defer file.Close()

    reader      := source.Read(file)
    lexer       := token.NewLexer(reader)
    resolver    := resolve.Run(lexer)
    interpreter := exec.NewInterpreter(resolver)


    exitCode, errs := interpreter.Run()

    for _, err := range errs {
        start, end := err.Location.Start, err.Location.End

        msg := start.AtPath(file.Name()) + ": " + err.Message + "\n"

        var maxCol int
        for i := start.Row; i <= end.Row; i++ {
            line := reader.GetBufferedLine(i - 1)
            if len(line) > maxCol {
                maxCol = len(line)
            }
            msg += fmt.Sprintf("%" + strconv.Itoa(codeColOffset) + "d", i) + " | " + line + "\n"
        }

        minCol := start.Col
        if end.Col < minCol {
            minCol = end.Col
        }

        if start.Row == end.Row {
            maxCol = end.Col
        }

        msg += strings.Repeat(" ", codeColOffset + 2 + minCol) + strings.Repeat("^", maxCol - minCol + 1)
        os.Stderr.WriteString(msg + "\n")
    }
    os.Exit(exitCode)
    return nil
}