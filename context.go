package requestid

import (
	"context"
	"errors"
)

type (
	requestIDKey struct{}
)

var (
	key = requestIDKey{}

	// ErrNotFound is the error when the context does not contain a reqestID.
	ErrNotFound = errors.New("context does not contain a requestID")
	// ErrWrongType is the error when the value in the context has the wrong type.
	ErrWrongType = errors.New("value is not of type string")
)

// Set sets the requestID to the context.
func Set(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, key, requestID)
}

// Get returns the stored value of the requestIDKey
// if no value is present (ErrNotFound) or for some reason
// the value is not a string (ErrWrongType) an error is returned.
func Get(ctx context.Context) (requestID string, err error) {
	val := ctx.Value(key)

	if val == nil {
		err = ErrNotFound

		return
	}

	var ok bool
	if requestID, ok = val.(string); !ok {
		err = ErrWrongType
	}

	return
}

// Copy gets the requestID from the source context, and copies it over
// to the target context.
func Copy(source context.Context, target *context.Context) error {
	requestID, err := Get(source)
	if err != nil {
		return err
	}

	*target = Set(*target, requestID)

	return nil
}
