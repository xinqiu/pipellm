package types

import "errors"

var (
	// ErrOpenAIResponse no response
	ErrOpenAIResponse = errors.New("no response")
	// ErrNoAPIKey no API key
	ErrNoAPIKey = errors.New("no API key")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
)
