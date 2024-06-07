package token

// defining a typedef kinda thing of C , to say that TokenType is an alias for type string 
type TokenType string


type Token struct{
    Type TokenType
    Literal String 
}

const {

    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // literals and idenfiers 
    IDENT = "IDENT"
    INT = "INT"

    // OPERATORS
    ASSIGN = "="
    PLUS = "+"

    LPAREN = ")"
}
