
# Writing an interpreter in go 

This book has the following chapters and we will go through them one by one . 

1  - Lexer or scanning
2  - Parser
3  - Evaluation
4  - Extending the interpreter

The book also adds REPL to the interpreter. REPL stands for Read Evaluate Print Loop and it's a simple interactive programming environment that takes user input, evaluates it, and returns the result to the user.

Let's Start with the first chapter which is Lexer or scanning.


## Lexer 

The Lexer basically converts the code you write into tokens. Tokens are the smallest unit of a program that is meaningful to the programming language.

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

Let's start going through functions now

- New() : This function initializes the lexer and returns a pointer to the lexer . This is the starting point of where this all compiler magic starts. you first dump your code and then just experiment with the lexer and see what tokens it generates .

- NextToken() : This function is the heart of the lexer . It reads the next character in the input and returns the corresponding token . It uses other functions like readChar() , peekChar() , isLetter() , isDigit() , readIdentifier() , readNumber() , skipWhitespace() to generate the token .

- readChar() : This function reads the next character in the input and advances the position in the input string . It also sets the position of the next character to read .

- peekChar() : This function returns the next character in the input string but does not advance the position in the input string . It is used to look ahead in the input string to see what the next character is . This function is mainly used to check if we are having a two character token or not like == , != etc .

- isLetter() : This function checks if the character is a letter or not . It is used to check if the character is a letter or not . If the character is a letter then it is part of an identifier .

- isDigit() : This function checks if the character is a digit or not . It is used to check if the character is a digit or not . If the character is a digit then it is part of a number .

- readIdentifier() : This function reads an identifier from the input string . It reads the characters until it encounters a non-letter character . It then sets the position to the next character to read .

- readNumber() : This function reads a number from the input string . It reads the characters until it encounters a non-digit character . It then sets the position to the next character to read .

- skipWhitespace() : This function skips the whitespace characters in the input string . It is used to skip the whitespace characters so that we can read the next token .

This is the basic overview of the lexer and the functions used in it . The Story gets interesting now as we dive into one of the most important parts of not only an interpreter but also computer science which is Parsing .


## Parser

Before entering to what functions are used in parser , let's understand what is parsing .


Parsing is the process of analyzing a string of symbols, either in natural language, computer languages or data structures, conforming to the rules of a formal grammar. In out case we will be parsing the tokens generated by the lexer into an Abstract Syntax Tree (AST) . We will divided the language into two parts 

1 - Expressions
2 - Statements

Expressions are the building blocks of a language and they produce a value . For example 5 + 5 is an expression which produces 10 . Whereas statements are the building blocks of a program and they do not produce a value . For example let x = 5 is a statement which does not produce a value . There are only two types of statements in our language which are let and return .


While parsing expressions, specifically operators , we have to take care of the precedence and associativity of the operators . For example in the expression 5 + 5 * 2 , we have to make sure that the multiplication operator is evaluated first and then the addition operator . This is because the multiplication operator has higher precedence than the addition operator . We also have to take care of the associativity of the operators . For example in the expression 5 - 5 - 2 , we have to make sure that the subtraction operator is evaluated from left to right . This is because the subtraction operator is left associative . 


There are multiple kinda of parsing techniques like Recursive Descent Parsing , Pratt Parsing etc . We will be using Pratt Parsing in our interpreter . Pratt Parsing is a top down operator precedence parsing technique which is used to parse expressions . It is a simple and efficient parsing technique which is easy to implement . It is also known as precedence climbing .


To explain pratt parsing , we will take an example of an expression 5 + 5 * 2 . The expression can be represented as a tree as follows :

```
    +
   / \
  5   *
     / \
    5   2
```

The tree is evaluated from the bottom to the top . The multiplication operator is evaluated first and then the addition operator . The tree is evaluated as follows :

```
    +
   / \
  5   10
```

There are all many pre built parsers available like yacc , bison etc . But we will be building our own parser in this book to understand the concepts of parsing .Understanding the concepts of parsing is very important as it is the heart of the interpreter . 


At last, before we enter into functions and all, lets discuss the importance of testing . Testing is very important in software development . It helps us to find bugs in our code and make sure that our code works as expected . We will be writing tests for our lexer and parser to make sure that they work We will adding _test file for every package we write such as lexer, parser, ast and token until now .


Now let's go through the functions and variables used in parser and understand them . Lets divide this into multiple groups and explain them one by one .

### Group 1 : AST

#### Variables  and there methods used in ast :

 - Node : This is the interface which is implemented by all the nodes in the AST . It has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .
 
 - Statement : This is the interface which is implemented by all the statement nodes in the AST . It has a statementNode() method which is used to differentiate between statement nodes and expression nodes . 

 - Expression : This is the interface which is implemented by all the expression nodes in the AST . It has a expressionNode() method which is used to differentiate between statement nodes and expression nodes .

 - Program : This is the root node of the AST . It has a Statements field which is a slice of Statement nodes . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

 - LetStatement : This is the node which represents the let statement in the AST . It has a Token field which is the token of the let statement . It also has a Name field which is the identifier of the let statement . It also has a Value field which is the expression of the let statement . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

 - ReturnStatement : This is the node which represents the return statement in the AST . It has a Token field which is the token of the return statement . It also has a ReturnValue field which is the expression of the return statement . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

 - ExpressionStatement : This is the node which represents the expression statement in the AST . It has a Token field which is the token of the expression statement . It also has a Expression field which is the expression of the expression statement . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

- PrefixExpression : This is the node which represents the prefix expression in the AST . It has a Token field which is the token of the prefix expression . It also has a Operator field which is the operator of the prefix expression . It also has a Right field which is the right expression of the prefix expression . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

- InfixExpression : This is the node which represents the infix expression in the AST . It has a Token field which is the token of the infix expression . It also has a Operator field which is the operator of the infix expression . It also has a Left field which is the left expression of the infix expression . It also has a Right field which is the right expression of the infix expression . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .


- Identifier : This is the node which represents the identifier in the AST . It has a Token field which is the token of the identifier . It also has a Value field which is the value of the identifier . It also has a TokenLiteral() method which returns the literal value of the token that the node represents . It is used to print the AST .

Now that we have understood the variables and methods used in the ast package , let's go through the functions used in the parser package .


### Group 2 : Parser 

In parser package, we start by creating a parser struct which has the following fields :

- l : This is a pointer to the lexer . It is used to read the tokens from the lexer .
- curToken : This is the current token that we are looking at . It is used to keep track of the current token .
- peekToken : This is the next token that we are going to look at . It is used to look ahead in the input string to see what the next token is .
- errors : This is a slice of strings which contains the errors that we encounter while parsing the input string . It is used to keep track of the errors that we encounter .
- prefixParseFns : This is a map which contains the prefix parsing functions for the different tokens . It is used to parse the prefix expressions .
- infixParseFns : This is a map which contains the infix parsing functions for the different tokens . It is used to parse the infix expressions .

Every method in this package is a method of the parser struct . The parser struct has the following methods :

- New() : This method initializes the parser and returns a pointer to the parser . It is the starting point of the parser . This is where we start parsing the input string . This method has following steps :

``` 
    P := &Parser{l: l , errors: []string{}} // Initialize the parser with the lexer and an empty slice of errors

    p.nextToken() // Read the first token from the lexer 
    p.nextToken() // Read the second token from the lexer 

    return p // Return the parser 

```
as we go through the parser, we will add more methods to the New() method . For now we are just initializing the parser and reading the first two tokens from the lexer . 


- ParseProgram() : This method parses the input string and returns the root node of the AST . It is the starting point of the parsing process . This method has following steps :

```
    program := &ast.Program{} // Create a new program node 

    for p.curToken.Type != token.EOF { // Loop until we reach the end of the input string 

        stmt := p.parseStatement() // Parse the statement 

        if stmt != nil { // If the statement is not nil 
            program.Statements = append(program.Statements, stmt) // Append the statement to the program 
        }

        p.nextToken() // Read the next token 

    }

    return program // Return the program 

```
- parseStatement() : This method parses the statement and returns the statement node of the AST . It has following steps :

```
    switch p.curToken.Type { // Switch on the current token type 

    case token.LET: // If the current token type is LET 
        return p.parseLetStatement() // Parse the let statement 

    case token.RETURN: // If the current token type is RETURN 
        return p.parseReturnStatement() // Parse the return statement 

    default: // If the current token type is not LET or RETURN 
        return p.parseExpressionStatement() // Parse the expression statement 

    }

```

As we said earlier, our language has only two types of statements which are let and return . If the current token type is LET, we parse the let statement . If the current token type is RETURN, we parse the return statement . If the current token type is not LET or RETURN, we parse the expression statement . 

- parseLetStatement() : This method parses the let statement and returns the let statement node of the AST . It has following steps :

```
    stmt := &ast.LetStatement{Token: p.curToken} // Create a new let statement node 

    if !p.expectPeek(token.IDENT) { // If the next token is not an identifier 
        return nil // Return nil 
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // Set the name of the let statement 

    if !p.expectPeek(token.ASSIGN) { // If the next token is not an assignment operator 
        return nil // Return nil 
    }

    p.nextToken() // Read the next token 

    stmt.Value = p.parseExpression(LOWEST) // Parse the expression 

    if p.peekTokenIs(token.SEMICOLON) { // If the next token is a semicolon 
        p.nextToken() // Read the next token 
    }

    return stmt // Return the let statement 

```


- parseReturnStatement() : This method parses the return statement and returns the return statement node of the AST . It has following steps :

```
    stmt := &ast.ReturnStatement{Token: p.curToken} // Create a new return statement node 

    p.nextToken() // Read the next token 

    stmt.ReturnValue = p.parseExpression(LOWEST) // Parse the expression 

    if p.peekTokenIs(token.SEMICOLON) { // If the next token is a semicolon 
        p.nextToken() // Read the next token 
    }

    return stmt // Return the return statement 

```

- parseExpressionStatement() : This method parses the expression statement and returns the expression statement node of the AST . It has following steps :

```
    stmt := &ast.ExpressionStatement{Token: p.curToken} // Create a new expression statement node 

    stmt.Expression = p.parseExpression(LOWEST) // Parse the expression 

    if p.peekTokenIs(token.SEMICOLON) { // If the next token is a semicolon 
        p.nextToken() // Read the next token 
    }

    return stmt // Return the expression statement 

```

The reason we are using "LOWEST" in parseExpression() is because we are not passing any precedence to the parseExpression() method . This is because we want to parse the entire expression without any precedence . We want to parse the entire expression as a single unit . we will deal with precedence in the parseExpression() method later on .

for precedence we create a map of precedence levels for the different operators . The precedence levels are as follows :

```
const (
    _ int = iota
    LOWEST
    EQUALS      // ==
    LESSGREATER // > or <
    SUM         // +
    PRODUCT     // *
    PREFIX      // -X or !X
    CALL        // myFunction(X)
)
```

Next up ,

- parseExpression() : This method parses the expression and returns the expression node of the AST . It has following steps :

```
    prefix := p.prefixParseFns[p.curToken.Type] // Get the prefix parsing function for the current token type 

    if prefix == nil { // If the prefix parsing function is nil 
        p.noPrefixParseFnError(p.curToken.Type) // Add an error to the parser 
        return nil // Return nil 
    }

    leftExp := prefix() // Parse the prefix expression 

    for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() { // Loop until we reach a semicolon or the precedence is less than the peek precedence 

        infix := p.infixParseFns[p.peekToken.Type] // Get the infix parsing function for the peek token type 

        if infix == nil { // If the infix parsing function is nil 
            return leftExp // Return the left expression 
        }

        p.nextToken() // Read the next token 

        leftExp = infix(leftExp) // Parse the infix expression 

    }

    return leftExp // Return the left expression 

```

Before dealing with infix and prefix expressions, we have to deal with errors and token methods . There are multiple methods for both tokens and error handling .

- peekError() : This method adds an error to the parser . It has following steps :

```
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", expected, p.peekToken.Type) // Create an error message 

    p.errors = append(p.errors, msg) // Append the error message to the errors slice 

```

- noPrefixParseFnError() : This method adds an error to the parser . It has following steps :

```
    msg := fmt.Sprintf("no prefix parse function for %s found", t) // Create an error message 

    p.errors = append(p.errors, msg) // Append the error message to the errors slice 

```

Now lets deal with token methods :

- peekTokenIs() : This method checks if the next token is of a certain type . It has following steps :

```
    return p.peekToken.Type == t // Return true if the next token is of the specified type 

```

- curTokenIs() : This method checks if the current token is of a certain type . It has following steps :

```
    return p.curToken.Type == t // Return true if the current token is of the specified type 

```

- expectPeek() : This method checks if the next token is of a certain type and advances the tokens . It has following steps :

```
    if p.peekTokenIs(t) { // If the next token is of the specified type 
        p.nextToken() // Read the next token 
        return true // Return true 
    } else { // If the next token is not of the specified type 
        p.peekError(t) // Add an error to the parser 
        return false // Return false 
    }

```

- nextToken() : This method advances the tokens . It has following steps :

```
    p.curToken = p.peekToken // Set the current token to the peek token 
    p.peekToken = p.l.NextToken() // Set the peek token to the next token from the lexer 

```

Now lets deal with Identifier and Number parsing  and then we can move to infix and prefix parsing .

- parseIdentifier() : This method parses the identifier and returns the identifier node of the AST . It has following steps :

```
    return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // Create a new identifier node 

```

- parseIntegerLiteral() : This method parses the integer literal and returns the integer literal node of the AST . It has following steps :

```
    lit := &ast.IntegerLiteral{Token: p.curToken} // Create a new integer literal node 

    value, err := strconv.ParseInt(p.curToken.Literal, 0, 64) // Parse the integer value 

    if err != nil { // If there is an error 
        msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal) // Create an error message 
        p.errors = append(p.errors, msg) // Append the error message to the errors slice 
        return nil // Return nil 
    }

    lit.Value = value // Set the value of the integer literal 

    return lit // Return the integer literal 

```


Now lets deal with infix and prefix parsing .

Before dealing with the functions lets understand the concept of infix and prefix parsing .

Infix parsing is the process of parsing an infix expression . An infix expression is an expression where the operator is in between the operands . For example 5 + 5 is an infix expression where the operator + is in between the operands 5 and 5 . Infix expressions are evaluated from left to right . Whereas prefix parsing is the process of parsing a prefix expression . A prefix expression is an expression where the operator is before the operands . For example -5 is a prefix expression where the operator - is before the operand 5 . Prefix expressions are evaluated from right to left . In our language we have only two prefix operators which are - and ! . We will be dealing with these operators in the parser .


Lets deal with prefix parsing first .


- registerPrefix() : This method registers a prefix parsing function for a token type . It has following steps :

```
    p.prefixParseFns[token.IDENT] = p.parseIdentifier // Register the parseIdentifier function for the IDENT token type 
    p.prefixParseFns[token.INT] = p.parseIntegerLiteral // Register the parseIntegerLiteral function for the INT token type 
    p.prefixParseFns[token.BANG] = p.parsePrefixExpression // Register the parsePrefixExpression function for the BANG token type 
    p.prefixParseFns[token.MINUS] = p.parsePrefixExpression // Register the parsePrefixExpression function for the MINUS token type 

```

we will add the above statements into the New() method of the parser . This is because we want to register the prefix parsing functions when we initialize the parser .


- parsePrefixExpression() : This method parses the prefix expression and returns the prefix expression node of the AST . It has following steps :

```
    expression := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal} // Create a new prefix expression node 

    p.nextToken() // Read the next token 

    expression.Right = p.parseExpression(PREFIX) // Parse the right expression 

    return expression // Return the prefix expression 

```

- parseInfixExpression() : This method parses the infix expression and returns the infix expression node of the AST . It has following steps :

```
    expression := &ast.InfixExpression{Token: p.curToken, Operator: p.curToken.Literal, Left: left} // Create a new infix expression node 

    precedence := p.curPrecedence() // Get the precedence of the current token 

    p.nextToken() // Read the next token 

    expression.Right = p.parseExpression(precedence) // Parse the right expression 

    return expression // Return the infix expression 

```

Lets deal with if statments now and then end with function literals and function calls .


- parseIfExpression() : This method parses the if expression and returns the if expression node of the AST . It has following steps :

```
    expression := &ast.IfExpression{Token: p.curToken} // Create a new if expression node 

    if !p.expectPeek(token.LPAREN) { // If the next token is not a left parenthesis 
        return nil // Return nil 
    }

    p.nextToken() // Read the next token 

    expression.Condition = p.parseExpression(LOWEST) // Parse the condition 

    if !p.expectPeek(token.RPAREN) { // If the next token is not a right parenthesis 
        return nil // Return nil 
    }

    if !p.expectPeek(token.LBRACE) { // If the next token is not a left brace 
        return nil // Return nil 
    }

    expression.Consequence = p.parseBlockStatement() // Parse the consequence 

    if p.peekTokenIs(token.ELSE) { // If the next token is ELSE 
        p.nextToken() // Read the next token 

        if !p.expectPeek(token.LBRACE) { // If the next token is not a left brace 
            return nil // Return nil 
        }

        expression.Alternative = p.parseBlockStatement() // Parse the alternative 

    }

    return expression // Return the if expression 

```

Lets under the above method more deeply as it is a bit complex . The function basically parses the if expression . The if expression has a condition, a consequence and an alternative . The condition is the expression that is evaluated to a boolean value . The consequence is the block of statements that is executed if the condition is true . The alternative is the block of statements that is executed if the condition is false . The if expression is represented as follows :

```
    if (condition) {
        consequence
    } else {
        alternative
    }

```
So, after taking the input string, we first create a new if expression node . We then parse the condition which is the expression inside the parentheses . We then parse the consequence which is the block of statements inside the left brace . We then check if there is an alternative which is the else block . If there is an alternative, we parse the alternative which is the block of statements inside the else block . We then return the if expression node .  

With that said , lets move to function literals and function calls .

- parseFunctionLiteral() : This method parses the function literal and returns the function literal node of the AST . It has following steps :

```
    lit := &ast.FunctionLiteral{Token: p.curToken} // Create a new function literal node 

    if !p.expectPeek(token.LPAREN) { // If the next token is not a left parenthesis 
        return nil // Return nil 
    }

    lit.Parameters = p.parseFunctionParameters() // Parse the function parameters 

    if !p.expectPeek(token.LBRACE) { // If the next token is not a left brace 
        return nil // Return nil 
    }

    lit.Body = p.parseBlockStatement() // Parse the function body 

    return lit // Return the function literal 

```

To understand this funciton, lets know how a function literal is represented in our language . A function literal is represented as follows :

```
    fn(x, y) {
        x + y
    }

```

A function literal has parameters and a body . The parameters are the identifiers inside the parentheses . The body is the block of statements inside the left brace . So, after taking the input string, we first create a new function literal node . We then parse the parameters which are the identifiers inside the parentheses . We then parse the body which is the block of statements inside the left brace . We then return the function literal node .

Function calls are represented as follows :

```
    add(5, 5)

```
Its pretty simple compared to function literals . Lets see the function now .

- parseCallExpression() : This method parses the call expression and returns the call expression node of the AST . It has following steps :

```
    exp := &ast.CallExpression{Token: p.curToken, Function: function} // Create a new call expression node 

    exp.Arguments = p.parseExpressionList(token.RPAREN) // Parse the arguments 

    return exp // Return the call expression 

```


Easy, it only has arguments and a function . We first create a new call expression node . We then parse the arguments which are the expressions inside the parentheses . We then return the call expression node . But we have to keep in mind that we have to add LPAREN to infixParseFns in the New() method of the parser . Because we have to parse the arguments inside the parentheses .

with the parsing done, lets make changes to our repl and get the parser working .

- We will add a new file repl.go in the main package . This file will contain the code for the REPL . The REPL is the interactive programming environment that takes user input, evaluates it, and returns the result to the user . The REPL has the following steps :

```
    l := lexer.New(input) // Create a new lexer with the input 
    p := parser.New(l) // Create a new parser with the lexer 

    program := p.ParseProgram() // Parse the program 

    if len(p.Errors()) != 0 { // If there are errors 
        printParserErrors(out, p.Errors()) // Print the errors 
        continue // Continue to the next iteration 
    }

    evaluator.Eval(program, env) // Evaluate the program 

    if len(p.Errors()) != 0 { // If there are errors 
        printParserErrors(out, p.Errors()) // Print the errors 
        continue // Continue to the next iteration 
    }

```

Thats it, Its Evaluation time now . Lets move to the next chapter and understand how evaluation works in an interpreter .


