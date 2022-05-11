package entity

import "errors"

var ErrInternalServer = errors.New("internal server error")

var ErrNotFoundAccount = errors.New("none account has been found")

var ErrInvalidLimit = errors.New("invalid limit")
