package repositoryports

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyTaken = errors.New("email already taken")
var ErrTransactionNotFound = errors.New("transaction not found")
