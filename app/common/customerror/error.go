package customerror

import "github.com/pkg/errors"

// ErrorNotFound returns when the object is not found
var ErrorNotFound = errors.New("not found")

// ErrorCpfOrSecretInvalid returns when username or secret is invalid
var ErrorCpfOrSecretInvalid = errors.New("CPF or secret invalid")

// ErrorNotEnoughAccountBalance returns when the account dont have enough balance to withdraw
var ErrorNotEnoughAccountBalance = errors.New("there is not enough account balance")

// ErrorCPFDuplicated returns when the CPF is already registered in DB
var ErrorCPFDuplicated = errors.New("CPF already in use")
