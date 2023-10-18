package auth

import (
	"github.com/golang-jwt/jwt"
)

type JwtTokenGenerator struct {
	secret     []byte
	expiration int
}

func NewJwtTokenGenerator(config Config) *JwtTokenGenerator {
	return &JwtTokenGenerator{config.GetSecret(), config.GetExpiration()}
}

func (g *JwtTokenGenerator) GenerateToken(claims map[string]interface{}) (string, error) {
	jwtClaims := jwt.MapClaims{}

	for key, value := range claims {
		jwtClaims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	return token.SignedString(g.secret)
}

func (g *JwtTokenGenerator) GetExpiration() int {
	return g.expiration
}
