package service

import "github.com/pkg/errors"

// ErrorNotFound returns when the object is not found
var ErrorNotFound = errors.New("not found")

// ErrorInvalidSecret returns when the secret is invalid
var ErrorInvalidSecret = errors.New("invalid secret")

// ErrorNotEnoughAccountBalance returns when the account dont have enough balance to withdraw
var ErrorNotEnoughAccountBalance = errors.New("there is not enough account balance")
