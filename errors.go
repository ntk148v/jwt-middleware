package jwt

import "errors"

var (
	ErrInvalidSigningMethod    = errors.New("JWT: invalid signing method")
	ErrNoHMACKey               = errors.New("JWT: no a HMAC key")
	ErrNoRSAKey                = errors.New("JWT: no a RSA key")
	ErrNoECKey                 = errors.New("JWT: no a EC key")
	ErrInvalidToken            = errors.New("JWT: invalid token")
	ErrGetTokenId              = errors.New("JWT: can not get id from token")
	ErrGetIssuedTime           = errors.New("JWT: can not get issued time from token")
	ErrGetData                 = errors.New("JWT: can not get data from token")
	ErrNoStore                 = errors.New("JWT: no store provided")
	ErrUnexpectedSigningMethod = errors.New("JWT: unexpected signing method")
	ErrTokenMalformed          = errors.New("JWT: token is malformed")
	ErrTokenNotActive          = errors.New("JWT: token is not valid yet")
	ErrTokenExpired            = errors.New("JWT: token is expired")
)
