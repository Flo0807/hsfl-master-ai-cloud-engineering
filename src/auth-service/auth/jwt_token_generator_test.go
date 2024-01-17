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
	generator := JwtTokenGenerator{privateKey}

	t.Run("generate valid token", func(t *testing.T) {
		// given
		givenClaims := map[string]interface{}{
			"sub":      123,
			"username": "user",
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
		assert.Equal(t, "user", claims["username"])
	})

	t.Run("validate valid token", func(t *testing.T) {
		token, err := generator.GenerateToken(map[string]interface{}{"test": "data"})

		if err != nil {
			t.Fatal(err)
		}

		claims, err := generator.ValidateToken(token)

		assert.NoError(t, err)
		assert.Equal(t, "data", claims["test"])
	})
}
