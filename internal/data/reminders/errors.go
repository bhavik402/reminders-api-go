package reminders

import (
	"errors"
)

var (
	ErrStorageNotSupported = errors.New("DB not Supported")
	ErrFailedToOpenDB = errors.New("failed to open sqlite DB")
	ErrFailedToOpenFile = errors.New("failed to open file")
	ErrFailedToReadFile = errors.New("failed to read file")
	ErrFailedToUnmarshalJson = errors.New("failed to unmarshal json")
)
