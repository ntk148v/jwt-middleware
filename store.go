package jwt

type Store interface {
	// Check checks whether a token has been revoked
	// If not, it will return some user data and nil.
	Check(tokenId string, issuedAt float64) (data map[string]interface{}, err error)
	// Revoke revokes a token which is no longer in use.
	// This case often happens when a user logs out.
	// or an authorization ends.
	Revoke(tokenId string) error
}
