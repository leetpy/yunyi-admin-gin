package config

type JwtConfig struct {
	AccessSignedKey  string `mapstructure:"access-signed-key"`
	RefreshSignedKey string `mapstructure:"refresh-signed-key"`
}
