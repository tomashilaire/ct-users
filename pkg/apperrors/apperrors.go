package apperrors

import "root/pkg/errors"

var (
	NotFound         = errors.Define("not_found")
	IllegalOperation = errors.Define("illegal_operation")
	InvalidInput     = errors.Define("invalid_input")
	Internal         = errors.Define("internal")
)
