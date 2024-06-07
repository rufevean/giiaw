# Lexer

 -  This is the first step of the interpreter
 -  we take the source code and convert it into tokes and then send it to AST ( abstract syntax tree)
 - a good lexer also attaches the line number, file number and coulumn number for future error recovery
 - for most languages whitespaces and new lines are insignificant ,same with out language so we dont tokenize themsame with out language. but in languages like python number of whitespaces is significant .
 -

## Getting started

- we first start by defining the token struct with token type and value
-  what we are going to do is create a 'nexttoken()'funciton which passes through every character and prints it token, its simple



