package sys

import (
	"github.com/pkg/errors"

	"github.com/nqxcode/platform_common/sys/codes"
)

type commonError struct {
	msg  string
	code codes.Code
}

// NewCommonError creates a new common error
func NewCommonError(msg string, code codes.Code) *commonError {
	return &commonError{msg, code}
}

// Error returns the error message
func (r *commonError) Error() string {
	return r.msg
}

// Code returns the error code
func (r *commonError) Code() codes.Code {
	return r.code
}

// IsCommonError checks if the error is a common error
func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

// GetCommonError returns the common error
func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
