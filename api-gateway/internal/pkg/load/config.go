package load

import "github.com/spf13/viper"

type Config struct {
	ApiGatewayHost string
	ApiGatewayPort int

	UserServiceHost string
	UserServicePort int

	MessageServiceHost string
	MessageServicePort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		ApiGatewayHost: viper.GetString("server.host"),
		ApiGatewayPort: viper.GetInt("server.port"),

		UserServiceHost: viper.GetString("user_service.host"),
		UserServicePort: viper.GetInt("user_service.port"),

		MessageServiceHost: viper.GetString("message_service.host"),
		MessageServicePort: viper.GetInt("message_service.port"),
	}
	return &conf, nil
}
