package cerr

import (
	"fmt"
)

// New returns a new CodedError.
func New() Error {
	return new(CodedError)
}

// CodedError is a standard implemented of Error.
type CodedError struct {
	// ErrCode contains an error code.
	ErrCode string
	// ErrInternal is an internal error.
	ErrInternal error
	// ErrShowInternalError defines whether or not to expose the internal error in any message output.
	ErrShowInternalError bool
}

// Code returns the error code.
func (e *CodedError) Code() string {
	return e.ErrCode
}

// Internal returns the errors internal message.
func (e *CodedError) Internal() error {
	return e.ErrInternal
}

// WithCode sets the errors code.
func (e *CodedError) WithCode(code string) Error {
	e.ErrCode = code
	return e
}

// WithInternal sets the errors internal error.
func (e *CodedError) WithInternal(err error) Error {
	e.ErrInternal = err
	return e
}

// ShowInternal causes Error() to include the internal error.
func (e *CodedError) ShowInternal() Error {
	e.ErrShowInternalError = true
	return e
}

// HideInternal causes Error() to exclude the internal error.
func (e *CodedError) HideInternal() Error {
	e.ErrShowInternalError = false
	return e
}

// Error returns the error message.
func (e *CodedError) Error() string {
	if e.ErrShowInternalError {
		return fmt.Sprintf("%s: %s", e.ErrCode, e.ErrInternal)
	} else {
		return fmt.Sprintf("%s", e.ErrCode)
	}
}

// Unwrap returns the internal error.
func (e *CodedError) Unwrap() error {
	return e.ErrInternal
}

// Is returns true if the target error is a CodedError either no Code
// or the same Code as e.
func (e *CodedError) Is(target error) bool {
	switch target := target.(type) {
	case *CodedError:
		targetCode := target.Code()
		return targetCode == "" || targetCode == e.ErrCode
	default:
		return false
	}
}
