package model

type (
	SvcConfig struct {
		URLGetRate             string `validate:"required"`
		URLGetRateMethod       string `validate:"required"`
		AccessKeyExchangeRates string `validate:"required"`
		RedisConfig            RedisConfig
		DatabaseConfig         DatabaseConfig
	}
	RedisConfig struct {
		Host    string `validate:"required"`
		KeyRate string `validate:"required"`
	}
	DatabaseConfig struct {
		Host     string `validate:"required"`
		User     string `validate:"required"`
		Password string `validate:"required"`
		Database string `validate:"required"`
	}
)
