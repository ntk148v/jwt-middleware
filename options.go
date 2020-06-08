package jwt

import (
	"time"
)

// Options is a struct for specifying configuration options
type Options struct {
	PrivateKeyLocation string
	PublicKeyLocation  string
	HMACKey            []byte
	SigningMethod      string
	TTL                time.Duration
	IsBearerToken      bool
	Header             string
	UserProperty       string
}
