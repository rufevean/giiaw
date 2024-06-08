# Writing an interpreter in go 

This book has the following chapters and we will go through them one by one . 

1  - Lexer or scanning
2  - Parser
3  - Evaluation
4  - Extending the interpreter

The book also adds REPL to the interpreter . REPL stands for Read Evaluate Print Loop and its a simple interactive programming environment that takes user input, evaluates it, and returns the result to the user.

Lets Start with the first chapter which is Lexer or scanning .


## Lexer 

The Lexer basically converts the code you write into tokens . Tokens are the smallest unit of a program that is meaningful to the programming language.

here is an example token.

```
{
  Type: token.INT,
  Literal: "5"
}
```

The above token is of type integer and its value is 5 . 

Lets understand about lexer more by noting the functions used here for and going through them one by one

- New()
- NextToken()
- readChar()
- peekChar()
- isLetter()
- isDigit()
- readIdentifier()
- readNumber()
- skipWhitespace()


Before working with the functions we have to tell the lexer what is a token and what is not . We will define the tokens in a const block in token.go file . 

an example will be :

```
package token 

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
   RETURN = "RETURN"
)
```

Here we have defined 3 tokens ILLEGAL , EOF and RETURN  , we will define more tokens as we go through the book .

Lets start going through funcitons now

- New() : This function initializes the lexer and returns a pointer to the lexer . This is the starting point of where this all compiler magic starts. you first dump your code and then just experiment with the lexer and see what tokens it generates .

- NextToken() : This function is the heart of the lexer . It reads the next character in the input and returns the corresponding token . It uses other functions like readChar() , peekChar() , isLetter() , isDigit() , readIdentifier() , readNumber() , skipWhitespace() to generate the token .

- readChar() : This function reads the next character in the input and advances the position in the input string . It also sets the position of the next character to read .

- peekChar() : This function returns the next character in the input string but does not advance the position in the input string . It is used to look ahead in the input string to see what the next character is . This function is mainly used to check if we are having a two character token or not like == , != etc .

- isLetter() : This function checks if the character is a letter or not . It is used to check if the character is a letter or not . If the character is a letter then it is part of an identifier .

- isDigit() : This function checks if the character is a digit or not . It is used to check if the character is a digit or not . If the character is a digit then it is part of a number .

- readIdentifier() : This function reads an identifier from the input string . It reads the characters until it encounters a non-letter character . It then sets the position to the next character to read .

- readNumber() : This function reads a number from the input string . It reads the characters until it encounters a non-digit character . It then sets the position to the next character to read .

- skipWhitespace() : This function skips the whitespace characters in the input string . It is used to skip the whitespace characters so that we can read the next token .

This is the basic overview of the lexer and the functions used in it . The Story gets interesting now as we dive into one of the most important part of not only an interpreter but also computer science which is Parsing .


