package service

import "github.com/pkg/errors"

// ErrorNotFound returns when the object is not found
var ErrorNotFound = errors.New("not found")

// ErrorInvalidSecret returns when the secret is invalid
var ErrorInvalidSecret = errors.New("invalid secret")
