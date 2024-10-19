package parser

import (
    "fmt"
    "strings"

    "github.com/mhdsulaimzed/py-go/lexer"
)

type ASTNode struct {
    Type     string
    Value    string
    Children []*ASTNode
}

func (n *ASTNode) String() string {
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("{Type: %s, Value: %s", n.Type, n.Value))
    if len(n.Children) > 0 {
        sb.WriteString(", Children: [")
        for i, child := range n.Children {
            if i > 0 {
                sb.WriteString(", ")
            }
            sb.WriteString(child.String())
        }
        sb.WriteString("]")
    }
    sb.WriteString("}")
    return sb.String()
}

func Parse(tokens []lexer.Token) *ASTNode {
    root := &ASTNode{Type: "Program", Children: []*ASTNode{}}
    i := 0

    for i < len(tokens) {
        if tokens[i].Type == lexer.DEF {
            funcNode := parseFunction(tokens, &i)
            root.Children = append(root.Children, funcNode)
        } else {
            i++
        }
    }

    return root
}

func parseFunction(tokens []lexer.Token, i *int) *ASTNode {
    funcNode := &ASTNode{Type: "FunctionDef", Children: []*ASTNode{}}
    *i++ 

    if *i < len(tokens) && tokens[*i].Type == lexer.IDENT {
        funcNode.Value = tokens[*i].Value
        *i++
    }

    
    *i += 3

    for *i < len(tokens) && tokens[*i].Type != lexer.DEF {
        if tokens[*i].Type == lexer.PRINT {
            printNode := parsePrint(tokens, i)
            funcNode.Children = append(funcNode.Children, printNode)
        } else {
            *i++
        }
    }

    return funcNode
}

func parsePrint(tokens []lexer.Token, i *int) *ASTNode {
    printNode := &ASTNode{Type: "Print", Children: []*ASTNode{}}
    *i++ 

    if *i < len(tokens) && tokens[*i].Type == lexer.LPAREN {
        *i++ 
    }

    if *i < len(tokens) && tokens[*i].Type == lexer.STRING {
        printNode.Children = append(printNode.Children, &ASTNode{Type: "String", Value: tokens[*i].Value})
        *i++
    }

    if *i < len(tokens) && tokens[*i].Type == lexer.RPAREN {
        *i++ 
    }

    return printNode
}