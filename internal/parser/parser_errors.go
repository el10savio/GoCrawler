package parser

import "errors"

var (
	ErrEmptyURL     = errors.New("empty URL provided")
	ErrInvalidURL   = errors.New("invalid URL provided")
	ErrEmptyURLHost = errors.New("empty URL host provided")
)
