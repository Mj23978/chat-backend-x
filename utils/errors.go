package utils

import "github.com/pkg/errors"

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok || cause.Cause() == nil {
			break
		}
		err = cause.Cause()
	}
	return err
}

// StatusCodeCarrier can be implemented by an error to support setting status codes in the error itself.
type StatusCodeCarrier interface {
	// StatusCode returns the status code of this error.
	StatusCode() int
}

// RequestIDCarrier can be implemented by an error to support error contexts.
type RequestIDCarrier interface {
	// RequestID returns the ID of the request that caused the error, if applicable.
	RequestID() string
}

// ReasonCarrier can be implemented by an error to support error contexts.
type ReasonCarrier interface {
	// Reason returns the reason for the error, if applicable.
	Reason() string
}

// DebugCarrier can be implemented by an error to support error contexts.
type DebugCarrier interface {
	// Debug returns debugging information for the error, if applicable.
	Debug() string
}

// StatusCarrier can be implemented by an error to support error contexts.
type StatusCarrier interface {
	// ID returns the error id, if applicable.
	Status() string
}

// DetailsCarrier can be implemented by an error to support error contexts.
type DetailsCarrier interface {
	// Details returns details on the error, if applicable.
	Details() map[string]interface{}
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}
