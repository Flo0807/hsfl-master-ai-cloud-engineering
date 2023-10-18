package auth

type TokenGenerator interface {
	GenerateToken(claims map[string]interface{}) (string, error)
	GetExpiration() int
}
