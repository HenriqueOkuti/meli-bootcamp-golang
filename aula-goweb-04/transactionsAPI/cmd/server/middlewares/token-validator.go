package token_validator

import "errors"

type TokenValidator struct {
	token string
}

type ITokenValidator interface {
	ValidateToken(token string, refToken string) error
}

func NewTokenValidator(token string) ITokenValidator {
	return &TokenValidator{
		token: token,
	}
}

func (t *TokenValidator) ValidateToken(token string, refToken string) error {
	if token != refToken {
		return errors.New("invalid token")
	}

	return nil
}
