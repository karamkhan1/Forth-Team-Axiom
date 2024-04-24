package Forth_Team_Axiom

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Constants to represent the different types of operations in the forth interpreter.
const (
	push_type = iota // auto-increment from zero for unique operation types
	add_type
	sub_type
	mul_type
	div_type
	mod_type
	equal_type
	less_type
	greater_type
	and_type
	or_type
	invert_type
	drop_type
	dup_type
	swap_type
	over_type
	rot_type
	print_type
	printstr_type
	emit_type
	cr_type
	quote_type
	colon_type
	semi_type
	word_type
	if_type
	else_type
	then_type
	do_type
	loop_type
	i_type
	variable_type
	constant_type
	at_type
	bang_type
	plus_bang_type
	question_type
	cells_type
	allot_type
	key_type
	oparen_type
	cparen_type
)

// SourceLocation holds the location in source code (for debugging)
type SourceLocation struct {
	fileName     string
	lineNumber   int
	columnNumber int
}

// LexicalToken represents a token with a value, type, and source location info.
type LexicalToken struct {
	content   string
	tokenType int
	location  SourceLocation
}

// NumberStack manages the integer stack operations.
type NumberStack struct {
	elements []int
	size     int
}

// ControlLoop manages the loop states to keep track of the iterations
type ControlLoop struct {
	index int
	limit int
}

// logError logs and halts execution if the interpreter ends up running into an error.
func logError(errorOccurred error) {
	if errorOccurred != nil {
		log.Fatal(errorOccurred)
	}
}

// addToStack pushes a value on top of the stack (push/add/ammend)
func (s *NumberStack) addToStack(value int) {
	s.elements = append(s.elements, value) // push the value onto the stack
	s.size++                               // update size
}

// removeFromStack pops and returns the top value from the stack (pop)
func (s *NumberStack) removeFromStack() int {
	if s.size == 0 {
		log.Fatal("stack underflow") // check for underflow
	}
	s.size--                  // update size
	return s.elements[s.size] // pop the value
}

// viewTopOfStack returns the top value of the stack without removing it (peek)
func (s *NumberStack) viewTopOfStack() int {
	if s.size == 0 {
		log.Fatal("stack underflow") // check for underflow
	}
	return s.elements[s.size-1] // return the last element
}

// isNumeric checks if a string is a numeric value
func isNumeric(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

// isWhitespace checks if a byte is a blank character
func isWhitespace(char byte) bool {
	switch char {
	case ' ', '\t', '\n': // simple switch checker
		return true
	default:
		return false
	}
}

// tokenizeFile converts the lines into a different objects of LexicalToken. it reads the whole file, breaks it into parts we can work with (tokens), and figures out what each part is (is it a number, an operation, etc.).
func tokenizeFile(filePath string) ([]LexicalToken, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err) // using fmt.Errorf to wrap and return an error with context
	}

	var tokens []LexicalToken // this will hold all our tokens
	var currentToken string
	line, col := 1, 1 // tracking line and column for error messages and debugging

	// helper function to add tokens to the list
	addToken := func(tokenType int) {
		if currentToken != "" {
			tokens = append(tokens, LexicalToken{
				content:   currentToken,
				tokenType: tokenType,
				location:  SourceLocation{fileName: filePath, lineNumber: line, columnNumber: col - len(currentToken)},
			})
			currentToken = "" // reset current token after adding
		}
	}

	for i, char := range string(fileContent) {
		switch {
		case isWhitespace(char):
			if char == '\n' {
				line++
				col = 1 // reset column at new line
			} else {
				col++
			}
			addToken(word_type) // end of a word, potentially, if we were building one
		case char == '+':
			addToken(word_type) // end previous token if it was building
			addToken(add_type)
		case char == '-':
			addToken(word_type) // end previous token if it was building
			addToken(sub_type)
		case char == '*':
			addToken(word_type) // end previous token if it was building
			addToken(mul_type)
		case char == '/':
			addToken(word_type) // end previous token if it was building
			addToken(div_type)
		case isNumeric(char):
			currentToken += string(char) // build the number
		default:
			currentToken += string(char) // build a word or operation name
		}
	}

	addToken(word_type) // check if there's one last token to add

	return tokens, nil // all good, return the tokens
}

// Interpreter stores the state and manages the execution of tokens
type Interpreter struct {
	numericStack      NumberStack
	definitions       map[string][]LexicalToken
	activeDefinition  bool
	currentDefinition string
	varMap            map[string]int
	constMap          map[string]int
	dataSegment       []int
}

// executeTokens processes and executes a series of tokens.
func (interp *Interpreter) executeTokens(tokens []LexicalToken) {
	// Processes and executes a series of tokens
}

// main initializes and starts the application.
func main() {
	// Bootstraps and starts the main application logic
}
