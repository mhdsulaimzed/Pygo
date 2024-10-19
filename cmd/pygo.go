// File: cmd/pygo.go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/mhdsulaimzed/py-go/lexer"
    "github.com/mhdsulaimzed/py-go/parser"
    "github.com/mhdsulaimzed/py-go/interpreter"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: pygo <filename>")
        return
    }

    fileName := os.Args[1]
    content, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    code := string(content)

    tokens := lexer.Lex(code)

    ast := parser.Parse(tokens)
	
    interpreter.Execute(ast)
}