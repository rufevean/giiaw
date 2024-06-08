
package token

// defining a typedef kinda thing of C , to say that TokenType is an alias for type string 
type TokenType string


type Token struct{
    Type TokenType
    Literal string 
}

const (
    // illegal signifies the token we dont know about 
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // literals and idenfiers 
    IDENT = "IDENT"
    INT = "INT"

    EQ = "=="
    NOT_EQ = "!="


    // OPERATORS
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LT = "<"
    GT = ">"

    LPAREN = ")"
    RPAREN = ")"
    LBRACE= "{"
    RBRACE= "}"
    
    SEMICOLON = ";"
    COMMA = ","

    // keywords

    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"   
    ELSE = "ELSE"
    RETURN = "RETURN"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}

