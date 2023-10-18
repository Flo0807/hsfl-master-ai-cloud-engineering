package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtTokenGenerator(t *testing.T) {
	generator := JwtTokenGenerator{[]byte("super_secret_key"), 3600}

	t.Run("generate valid token", func(t *testing.T) {
		// given
		claims := map[string]interface{}{
			"sub":   123,
			"email": "email@example.com",
		}

		// when
		token, err := generator.GenerateToken(claims)

		// test
		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQGV4YW1wbGUuY29tIiwic3ViIjoxMjN9.LwRCkXk83at4OLC-3NKSfErtEmlNIH5E7ga3T_U7qR8", token)
	})
}
