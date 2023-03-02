package types

import "errors"

var (
	// ErrOpenAIResponse no response
	ErrOpenAIResponse = errors.New("no response")
	// ErrNoAPIKey no API key
	ErrNoAPIKey = errors.New("no API key")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrOnlySupportOneArg `Run` supports only one positional argument
	ErrOnlySupportOneArg = errors.New("`Run` supports only one positional argument")
	// ErrTextLoader TextLoader read file error
	ErrTextLoader = errors.New("TextLoader read file error")
)
