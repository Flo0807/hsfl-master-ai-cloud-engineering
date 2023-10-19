package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtTokenGenerator(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	generator := JwtTokenGenerator{privateKey, 3600}

	t.Run("generate valid token", func(t *testing.T) {
		// given
		givenClaims := map[string]interface{}{
			"sub":   123,
			"email": "email@example.com",
		}

		// when
		token, err := generator.GenerateToken(givenClaims)

		// test
		assert.NoError(t, err)
		tokenParts := strings.Split(token, ".")
		assert.Len(t, tokenParts, 3)

		b, _ := base64.
			StdEncoding.
			WithPadding(base64.NoPadding).
			DecodeString(tokenParts[1])

		var claims map[string]interface{}
		json.Unmarshal(b, &claims)

		assert.Equal(t, float64(123), claims["sub"])
		assert.Equal(t, "email@example.com", claims["email"])
	})
}
