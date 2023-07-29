package configs

import (
	"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"
)

type conf struct {
	APIName    string `mapstructure:"API_NAME" valid:"required"`
	APIPort    int    `mapstructure:"API_PORT" valid:"required"`
	APIEnv     string `mapstructure:"API_ENV" valid:"required"`
	DBHost     string `mapstructure:"DB_HOST" valid:"required"`
	DBPort     int    `mapstructure:"DB_PORT" valid:"required"`
	DBUser     string `mapstructure:"DB_USER" valid:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" valid:"required"`
	DBName     string `mapstructure:"DB_NAME" valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
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

	if err := cfg.isValid(); err != nil {
		panic(err)
	}

	return cfg, nil
}

func (c *conf) isValid() error {
	_, err := govalidator.ValidateStruct(c)

	if err != nil {
		return err
	}

	return nil
}
