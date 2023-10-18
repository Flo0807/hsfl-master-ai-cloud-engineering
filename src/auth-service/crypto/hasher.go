package crypto

type Hasher interface {
	Hash([]byte) ([]byte, error)
	Validate([]byte, []byte) bool
}
