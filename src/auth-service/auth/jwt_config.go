package auth

type JwtConfig struct {
	Secret      string           `yaml:"secret"`
	AccessToken ExpirationConfig `yaml:"access_token"`
}

type ExpirationConfig struct {
	Expiration int `yaml:"expiration"`
}

func (c JwtConfig) GetSecret() []byte {
	return []byte(c.Secret)
}

func (c JwtConfig) GetExpiration() int {
	return c.AccessToken.Expiration
}
