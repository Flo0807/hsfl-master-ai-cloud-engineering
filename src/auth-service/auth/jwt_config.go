package auth

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

type JwtConfig struct {
	SignKey string `yaml:"signKey"`
}

func (c JwtConfig) GetPrivateKey() (any, error) {
	bytes, err := os.ReadFile(c.SignKey)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	return x509.ParseECPrivateKey(block.Bytes)
}
