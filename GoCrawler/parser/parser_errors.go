package parser

import "errors"

var (
	// ErrEmptyURL Error in the case empty URL provided
	ErrEmptyURL = errors.New("empty URL provided")
	// ErrInvalidURL Error in the case invalid URL provided
	ErrInvalidURL = errors.New("invalid URL provided")
	// ErrEmptyURLHost Error in the case empty URL host provided
	ErrEmptyURLHost = errors.New("empty URL host provided")
)
