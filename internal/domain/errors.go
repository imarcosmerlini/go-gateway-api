package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicatedAPIKey   = errors.New("api key already exists")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized acess")
)
