package interpreter

import (
    "fmt"
    "strings"

    "github.com/mhdsulaimzed/py-go/parser"
)

func Execute(ast *parser.ASTNode) {
    for _, node := range ast.Children {
        if node.Type == "FunctionDef" && node.Value == "main" {
            executeFunction(node)
        }
    }
}

func executeFunction(node *parser.ASTNode) {
    for _, stmt := range node.Children {
        if stmt.Type == "Print" {
            executePrint(stmt)
        }
    }
}

func executePrint(node *parser.ASTNode) {
    if len(node.Children) > 0 && node.Children[0].Type == "String" {
        value := strings.Trim(node.Children[0].Value, "\"")
        fmt.Println(value)
    }
}