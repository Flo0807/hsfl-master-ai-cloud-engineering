package auth

type Config interface {
	GetSecret() []byte
	GetExpiration() int
}
