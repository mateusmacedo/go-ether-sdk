package err

import "errors"

var (
	ErrHandlerNotImplemented = errors.New("Handler not implemented")
	ErrMessageNotSupported   = errors.New("message type not supported")
)