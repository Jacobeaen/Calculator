package rpn

import "errors"

var (
	ErrForbiddenSymbols        = errors.New("string contains forbiden symbols")
	ErrBracketsPair            = errors.New("one of bracket hasn't pair")
	ErrOperationWithoutNumbers = errors.New("can't do operation. there is no numbers")
	ErrZeroDivison             = errors.New("zero division error")
	ErrUnknownOperation        = errors.New("unknow operation")
	ErrEmptyExpression         = errors.New("empty expression")
	ErrBrackets                = errors.New("brackets error")
)
