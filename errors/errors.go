package errors

import (
	"fmt"
	"github.com/vistormu/go-dsa/ansi"
)

type ErrorType string

type Error struct {
	Type     ErrorType
	message  string
	metadata map[string]any
	err      error
}

func New(errorType ErrorType) *Error {
	return &Error{
		Type:     errorType,
		message:  fmt.Sprintf("%s-> %s%s\n", ansi.Red, errorType, ansi.Reset),
		metadata: make(map[string]any),
	}
}

func (e *Error) Error() string {
	if e.metadata != nil {
		for k, v := range e.metadata {
			e.message += fmt.Sprintf("   |> %s: %v\n", k, v)
		}
	}

	if e.err != nil {
		e.message += fmt.Sprintf("   |> full error: %v\n", e.err)
	}

	return e.message
}

func (e *Error) With(metadata ...any) *Error {
	if len(metadata)%2 != 0 {
		panic("metadata must be in key-value pairs")
	}

	for i := 0; i < len(metadata); i += 2 {
		key, ok := metadata[i].(string)
		if !ok {
			panic("key must be a string")
		}
		value := metadata[i+1]
		e.metadata[key] = value
	}

	return e
}

func (e *Error) Wrap(err error) *Error {
	e.err = err
	return e
}

func (e *Error) Print() {
	fmt.Println(e.Error())
}
