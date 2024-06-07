package token

// defining a typedef kinda thing of C , to say that TokenType is an alias for type string 
type TokenType string


type Token struct{
    Type TokenType
    Literal String 
}

const {
    // illegal signifies the token we dont know about 
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // literals and idenfiers 
    IDENT = "IDENT"
    INT = "INT"

    // OPERATORS
    ASSIGN = "="
    PLUS = "+"

    LPAREN = ")"
    RPAREN = ")"
    LBRACE= "{"
    RBRACE= "}"


    // keywords

    FUNCTION = "FUNCTION"
    LET = "LET"
}


