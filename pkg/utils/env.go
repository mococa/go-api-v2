package utils

import "github.com/spf13/viper"

type Env struct {
	DB_URI string `mapstructure:"DB_URI"`
}

func LoadConfig() (config Env, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("server")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	return
}
