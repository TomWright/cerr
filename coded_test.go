package cerr_test

import (
	"errors"
	"fmt"
	"github.com/tomwright/cerr"
	"testing"
)

// TestCodedError_FactoryAndSetters ensures the setters of a CodedError work as expected.
func TestCodedError_FactoryAndSetters(t *testing.T) {
	t.Parallel()

	someInternalError := errors.New("name is missing")
	errCode := "INVALID_NAME"

	e := cerr.New()

	t.Run("Code", func(t *testing.T) {
		e = e.WithCode(errCode)
		if exp, got := errCode, e.Code(); exp != got {
			t.Errorf("expected error code `%s`, got `%s`", exp, got)
		}
	})
	t.Run("Internal", func(t *testing.T) {
		e = e.WithInternal(someInternalError)
		if internalErr := e.Internal(); internalErr != someInternalError {
			t.Errorf("unexpected internal error: %v", internalErr)
		}
	})
	t.Run("ShowHideInternal", func(t *testing.T) {
		if show := e.(*cerr.CodedError).ErrShowInternalError; show != false {
			t.Error("expected to hide internal error")
		}
		e = e.ShowInternal()
		if show := e.(*cerr.CodedError).ErrShowInternalError; show != true {
			t.Error("expected to show internal error")
		}
		e = e.HideInternal()
		if show := e.(*cerr.CodedError).ErrShowInternalError; show != false {
			t.Error("expected to hide internal error")
		}
	})
}

// TestCodedError_As ensures errors.Is returns the correct response with different errors.
func TestCodedError_Is(t *testing.T) {
	t.Parallel()

	someInternalError := errors.New("name is missing")
	errCode := "INVALID_NAME"

	e := cerr.New().
		WithCode(errCode).
		WithInternal(someInternalError)

	t.Run("CodedError", func(t *testing.T) {
		if !errors.Is(e, e) {
			t.Error("error is not the exact same CodedError")
		}

		if !errors.Is(e, &cerr.CodedError{}) {
			t.Error("error is not a CodedError")
		}

		if !errors.Is(e, cerr.New().WithCode(errCode)) {
			t.Error("error is not a CodedError with the same code")
		}

		if errors.Is(e, cerr.New().WithCode("DIFFERENT_CODE")) {
			t.Error("error should not a CodedError with a different code")
		}
	})

	t.Run("InternalError", func(t *testing.T) {
		if !errors.Is(e, someInternalError) {
			t.Error("error is not the correct internal error")
		}

		if errors.Is(e, errors.New("different error")) {
			t.Error("error should not be a different internal error")
		}
	})

	t.Run("NonError", func(t *testing.T) {
		if errors.Is(e, nil) {
			t.Error("error should not be a nil")
		}
	})
}

// TestCodedError_As ensures errors.As returns the correct response with different errors.
func TestCodedError_As(t *testing.T) {
	t.Parallel()

	someInternalError := errors.New("name is missing")
	errCode := "INVALID_NAME"

	e := cerr.New().
		WithCode(errCode).
		WithInternal(someInternalError)

	t.Run("CodedError", func(t *testing.T) {
		var outE *cerr.CodedError
		if !errors.As(e, &outE) {
			t.Error("error is not a CodedError")
			return
		}

		if exp, got := errCode, outE.ErrCode; exp != got {
			t.Errorf("expected error code `%s`, got `%s`", exp, got)
		}
		if internalErr := outE.ErrInternal; internalErr != someInternalError {
			t.Errorf("unexpected internal error: %v", internalErr)
		}
		if show := outE.ErrShowInternalError; show != false {
			t.Error("expected to hide internal error")
		}
	})
}

// TestCodedError_Unwrap ensures that the correct error is unwrapped.
func TestCodedError_Unwrap(t *testing.T) {
	t.Parallel()

	someInternalError := errors.New("name is missing")
	errCode := "INVALID_NAME"

	e := cerr.New().
		WithCode(errCode).
		WithInternal(someInternalError)

	if e := errors.Unwrap(e); e != someInternalError {
		t.Errorf("unexpected unwrapped error: %v", e)
	}
}

// TestCodedError_Error ensures the correct error messages are returned.
func TestCodedError_Error(t *testing.T) {
	t.Parallel()

	someInternalError := errors.New("name is missing")
	errCode := "INVALID_NAME"

	e := cerr.New().
		WithCode(errCode).
		WithInternal(someInternalError)

	t.Run("ShowInternal", func(t *testing.T) {
		e = e.ShowInternal()
		if exp, got := fmt.Sprintf("%s: %s", errCode, someInternalError), e.Error(); exp != got {
			t.Errorf("expected error message `%s`, got `%s`", exp, got)
		}
	})

	t.Run("HideInternal", func(t *testing.T) {
		e = e.HideInternal()
		if exp, got := fmt.Sprintf("%s", errCode), e.Error(); exp != got {
			t.Errorf("expected error message `%s`, got `%s`", exp, got)
		}
	})
}
