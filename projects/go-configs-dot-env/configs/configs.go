package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	APIName    string `mapstructure:"API_NAME"`
	APIPort    int    `mapstructure:"API_PORT"`
	APIEnv     string `mapstructure:"API_ENV"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
}

func InitConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}

func (c conf) Validate() bool {
	var isValid bool = true

	if c.APIName == "" {
		isValid = false
	}

	return isValid
}
