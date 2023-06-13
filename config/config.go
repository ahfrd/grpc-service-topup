package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"MYPORT"`
	DBUrl         string `mapstructure:"DB_URL"`
	DB            string `mapstructure:"DB"`
	EmoneySvcUrl  string `mapstructure:"EMONEY_SVC_URL"`
	HistorySvcUrl string `mapstructure:"HISTORY_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
