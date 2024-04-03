package err

import "errors"

var (
	ErrHandlerNotImplemented = errors.New("Handler not implemented")
	ErrMessageNotSupported   = errors.New("message type not supported")
	ErrMessageContentEmpty   = errors.New("message content is empty")
	ErrInvalidMessageContent = errors.New("message content is invalid")
	ErrInternalHandler       = errors.New("internal handler error")
	ErrPersistenceConflict  = errors.New("persistence conflict")
)