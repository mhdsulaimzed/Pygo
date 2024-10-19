package lexer

import (
    "strings"
    "unicode"
)

type TokenType string

const (
    DEF    TokenType = "DEF"
    IDENT  TokenType = "IDENT"
    LPAREN TokenType = "LPAREN"
    RPAREN TokenType = "RPAREN"
    COLON  TokenType = "COLON"
    STRING TokenType = "STRING"
    PRINT  TokenType = "PRINT"
)

type Token struct {
    Type  TokenType
    Value string
}

func Lex(input string) []Token {
    var tokens []Token
    lines := strings.Split(input, "\n")

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        for i := 0; i < len(line); i++ {
            switch {
            case strings.HasPrefix(line[i:], "def"):
                tokens = append(tokens, Token{DEF, "def"})
                i += 2
            case strings.HasPrefix(line[i:], "print"):
                tokens = append(tokens, Token{PRINT, "print"})
                i += 4
            case line[i] == '(':
                tokens = append(tokens, Token{LPAREN, "("})
            case line[i] == ')':
                tokens = append(tokens, Token{RPAREN, ")"})
            case line[i] == ':':
                tokens = append(tokens, Token{COLON, ":"})
            case line[i] == '"':
                end := strings.Index(line[i+1:], "\"")
                if end != -1 {
                    tokens = append(tokens, Token{STRING, line[i : i+end+2]})
                    i += end + 1
                }
            case unicode.IsLetter(rune(line[i])):
                end := i
                for end < len(line) && (unicode.IsLetter(rune(line[end])) || unicode.IsDigit(rune(line[end]))) {
                    end++
                }
                tokens = append(tokens, Token{IDENT, line[i:end]})
                i = end - 1
            }
        }
    }

    return tokens
}