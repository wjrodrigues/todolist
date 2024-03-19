package configs

import (
	"github.com/spf13/viper"
)

var values *env

type env struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	Secret        string `mapstructure:"SECRET"`
}

func LoadEnv(path, file string) (env, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return env{}, err
	}

	err = viper.Unmarshal(&values)

	return *values, err
}
