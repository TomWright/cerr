package cerr

// Error defines an error that has a error code, and with a visible or hidden internal error.
type Error interface {
	// Code returns the error code.
	Code() string
	// Internal returns the errors internal message.
	Internal() error

	// WithCode sets the errors code.
	WithCode(code string) Error
	// WithInternal sets the errors internal error.
	WithInternal(err error) Error

	// ShowInternal causes Error() to include the internal error.
	ShowInternal() Error
	// HideInternal causes Error() to exclude the internal error.
	HideInternal() Error

	// Error implements the Error interface.
	Error() string
}
