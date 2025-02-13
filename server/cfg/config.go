package cfg

import "github.com/spf13/viper"

type Configuration struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	WebserverPort string `mapstructure:"WEBSERVER_PORT"`
}

func LoadConfiguration(path string) (*Configuration, error) {

	var configuration Configuration

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
