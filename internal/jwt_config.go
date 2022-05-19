package internal

type JWTConfig struct {
	SingingKey string `mapstructure:"key" json:"key"`
}
