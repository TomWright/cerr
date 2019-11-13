package cerr_test

import (
	"errors"
	"fmt"
	"github.com/tomwright/cerr"
)

// ExampleNew demonstrates CodedError usage.
func ExampleNew() {
	// someErrorCode is an error code that we can show to an end user without concern of exposing the
	// internal workings of our system or any potentially secret values.
	// This could also be used as part of a translation key to support multiple languages.
	someErrorCode := "InternalServerError"
	// someInternalError would be an internal error that we do not want to display to an end user
	// but we do want to print out in logs for example.
	someInternalError := errors.New("could not execute sql query: database connection timed out")

	// define an error with the above code and error
	var err error = cerr.New().
		WithCode(someErrorCode).
		WithInternal(someInternalError)

	// by default, the internal error is not displayed in the error message.
	// this would be used to return an error to a http endpoint for example.
	fmt.Println("write error to http:", err.Error())

	// now we may want to log the error so as it will displayed in stdout
	var codedErr cerr.Error
	if errors.As(err, &codedErr) {
		// we would want the internal message in logs so as we can debug any issues
		err = codedErr.ShowInternal()
	}
	fmt.Println("write error to stdout:", err.Error())

	// Output:
	// write error to http: InternalServerError
	// write error to stdout: InternalServerError: could not execute sql query: database connection timed out
}

// ExampleCodedError_Is Error comparison.
func ExampleCodedError_Is() {
	// fairly generic error codes to be sent back to a client
	errCodeInternalServer := "InternalServerError"
	errCodeUserNotFound := "UserNotFound"

	// specific internal error messages to be included in logs
	errUsernameNotFound := errors.New("the given username does not exist in the database")
	errInvalidPassword := errors.New("the given password does not match")

	// define an error that shows a user could not be found because the given username didn't exist in the database
	var err error = cerr.New().
		WithCode(errCodeUserNotFound).
		WithInternal(errUsernameNotFound)

	fmt.Println("is coded:", errors.Is(err, cerr.New()))
	fmt.Println("is internal server error:", errors.Is(err, cerr.New().WithCode(errCodeInternalServer)))
	fmt.Println("is user not found:", errors.Is(err, cerr.New().WithCode(errCodeUserNotFound)))
	fmt.Println("is username not found in db:", errors.Is(err, errUsernameNotFound))
	fmt.Println("is invalid password:", errors.Is(err, errInvalidPassword))

	// Output:
	// is coded: true
	// is internal server error: false
	// is user not found: true
	// is username not found in db: true
	// is invalid password: false
}
