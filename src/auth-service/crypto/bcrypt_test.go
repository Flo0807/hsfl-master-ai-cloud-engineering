package crypto

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptHasher(t *testing.T) {
	hasher := NewBcryptHasher()

	t.Run("hash", func(t *testing.T) {
		// given
		data := []byte("testpassword")

		// when
		hash, err := hasher.Hash(data)

		// then
		assert.NoError(t, err)
		assert.Len(t, hash, 60)
		assert.Regexp(t, regexp.MustCompile(`\$2a\$10\$(.*)`), string(hash))
	})

	t.Run("validate", func(t *testing.T) {
		// given
		data := []byte("testpassword")
		hash := []byte("$2a$10$p0AnxxB2iEvIUMHW2CQKceXSgK1kiGoGCsnOLsuq/JYTMHIM7W0ly")

		// when
		valid := hasher.Validate(data, hash)

		// then
		assert.True(t, valid)
	})

	t.Run("validate wrong hash", func(t *testing.T) {
		// given
		data := []byte("testpassword")
		hash := []byte("$2a$10$p0AnxxB2iEvIUMHW2CQKceXSgK1kiGoGCsnOLsuq/JYTMHIM7W0lx")

		// when
		valid := hasher.Validate(data, hash)

		// then
		assert.False(t, valid)
	})
}
