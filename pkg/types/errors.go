package types

import "errors"

var (
	ErrOpenAIResponse = errors.New("no response")
	ErrNoAPIKey       = errors.New("no API key")
	ErrNotImplemented = errors.New("not implemented")
)
