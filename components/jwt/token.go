package jwt

type TokenProvider interface {
	GenerateAccessToken(payload Payload, expiry int) (*string, error)
	Decode(token string) (*Payload, error)
}
