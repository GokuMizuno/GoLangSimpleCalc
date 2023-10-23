package ComplexCalc

import "fmt"

const (
	TokenTypeOperand  = 1
	TokenTypeOperator = 2
)

// Token represents an abstract token object in RPN(Reverse Polish notation) which could either be an operator or operand.
type Token struct {
	Type  int
	Value interface{}
}

// NewRPNOperandToken creates an instance of operand Token with specified value.
func NewRPNOperandToken(val int) *Token {
	return NewToken(val, TokenTypeOperand)
}

// NewRPNOperatorToken creates an instance of operator Token with specified value.
func NewRPNOperatorToken(val string) *Token {
	return NewToken(val, TokenTypeOperator)
}

// NewToken creates an instance of Token with specified value and type.
func NewToken(val interface{}, typ int) *Token {
	return &Token{Value: val, Type: typ}
}

// IsOperand determines whether a token is an operand with a specified value.
func (token *Token) IsOperand(val int) bool {
	return token.Type == TokenTypeOperand && token.Value.(int) == val
}

// IsOperator determines whether a token is an operator with a specified value.
func (token *Token) IsOperator(val string) bool {
	return token.Type == TokenTypeOperator && token.Value.(string) == val
}

// GetDescription returns a string that describes the token.
func (token *Token) GetDescription() string {
	return fmt.Sprintf("(%d)%v", token.Type, token.Value)
}
