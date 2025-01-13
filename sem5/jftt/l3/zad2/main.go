package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Characteristic = 1234577

type Parser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
	errors       []string
	rpn          []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
		rpn:    []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Parse() (string, int64, error) {
	result, err := p.parseExpression(0)
	if err != nil {
		return "", 0, err
	}

	if p.currentToken.Type != EOL && p.currentToken.Type != EOF {
		return "", 0, fmt.Errorf("unexpected token after expression")
	}

	rpn := strings.Join(p.rpn, " ")
	return rpn, result, nil
}

const (
	LOWEST     int = iota
	ADD_SUB_OP     // + -
	MUL_DIV_OP     // * /
	POW_OP         // ^
	UNARY_OP       // -X
)

// getPrecedence returns the precedence of the current operator
func (p *Parser) getPrecedence() int {
	switch p.currentToken.Type {
	case ADD, SUB:
		return ADD_SUB_OP
	case MUL, DIV:
		return MUL_DIV_OP
	case POW:
		return POW_OP
	default:
		return LOWEST
	}
}

// parseExpression parses expressions based on precedence
func (p *Parser) parseExpression(precedence int) (int64, error) {
	var leftVal int64
	var err error

	switch p.currentToken.Type {
	case NUM:
		leftVal, err = p.parseNumber()
		if err != nil {
			return 0, err
		}
	case SUB:
		// Handle unary negation
		p.nextToken()
		value, err := p.parseExpression(UNARY_OP)
		if err != nil {
			return 0, err
		}
		leftVal = FieldSub(0, value)
		p.rpn = append(p.rpn, strconv.FormatInt(leftVal, 10))
	case LPAREN:
		p.nextToken()
		leftVal, err = p.parseExpression(LOWEST)
		if err != nil {
			return 0, err
		}
		if p.currentToken.Type != RPAREN {
			return 0, fmt.Errorf("expected ')'")
		}
		p.nextToken()
	default:
		return 0, fmt.Errorf("unexpected token: %s", p.currentToken.Value)
	}

	for p.currentToken.Type != EOL && p.currentToken.Type != EOF && precedence < p.getPrecedence() {
		op := p.currentToken.Type
		p.nextToken()
		rightVal, err := p.parseExpression(p.getPrecedence())
		if err != nil {
			return 0, err
		}

		switch op {
		case ADD:
			leftVal = FieldAdd(leftVal, rightVal)
			p.rpn = append(p.rpn, "+")
		case SUB:
			leftVal = FieldSub(leftVal, rightVal)
			p.rpn = append(p.rpn, "-")
		case MUL:
			leftVal = FieldMul(leftVal, rightVal)
			p.rpn = append(p.rpn, "*")
		case DIV:
			if rightVal == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			leftVal = FieldDiv(leftVal, rightVal)
			p.rpn = append(p.rpn, "/")
		case POW:
			leftVal = FieldPow(leftVal, rightVal)
			p.rpn = append(p.rpn, "^")
		default:
			return 0, fmt.Errorf("unknown operator")
		}
	}

	return leftVal, nil
}

// parseNumber parses a number token
func (p *Parser) parseNumber() (int64, error) {
	val, err := strconv.ParseInt(p.currentToken.Value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", p.currentToken.Value)
	}
	val = val % Characteristic
	p.rpn = append(p.rpn, strconv.FormatInt(val, 10))
	p.nextToken()
	return val, nil
}

// Field Arithmetic Operations

// FieldAdd performs addition in the finite field
func FieldAdd(a, b int64) int64 {
	return (a + b) % Characteristic
}

// FieldSub performs subtraction in the finite field
func FieldSub(a, b int64) int64 {
	return (a - b + Characteristic) % Characteristic
}

// FieldMul performs multiplication in the finite field
func FieldMul(a, b int64) int64 {
	return (a * b) % Characteristic
}

// FieldPow performs exponentiation in the finite field
func FieldPow(a, b int64) int64 {
	if b < 0 {
		// Handle negative exponents by converting to positive using Fermat's little theorem
		b = (Characteristic - 1) + b
	}
	result := int64(1)
	base := a % Characteristic
	exponent := b

	for exponent > 0 {
		if exponent%2 == 1 {
			result = FieldMul(result, base)
		}
		base = FieldMul(base, base)
		exponent /= 2
	}
	return result
}

// FieldDiv performs division in the finite field
func FieldDiv(a, b int64) int64 {
	if b == 0 {
		// Division by zero should be handled by the caller
		return 0
	}
	inv := FieldPow(b, Characteristic-2)
	return FieldMul(a, inv)
}

// Main Function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var currentLine strings.Builder

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		// Handle line continuation
		if strings.HasSuffix(line, "\\") {
			currentLine.WriteString(strings.TrimSuffix(line, "\\"))
			continue
		} else {
			currentLine.WriteString(line)
		}

		input := currentLine.String()
		currentLine.Reset()

		// Trim spaces
		input = strings.TrimSpace(input)

		// Skip empty lines
		if input == "" {
			continue
		}

		// Initialize Lexer and Parser
		lexer := NewLexer(input)
		parser := NewParser(lexer)

		// Parse and evaluate
		rpn, result, err := parser.Parse()
		if err != nil {
			fmt.Println("Błąd.")
			continue
		}

		// Print RPN and result
		if rpn != "" {
			fmt.Println(rpn)
		}
		fmt.Printf("Wynik: %d\n", result)
	}

	// Handle any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
