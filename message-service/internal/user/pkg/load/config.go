package load

import "github.com/spf13/viper"

type Mongosh struct {
	MongoHost string
	MongoPort int
	// MongoUser       string
	// MongoPassword   string
	MongoDatabase   string
	MongoCollection string
}

type Config struct {
	Mongosh Mongosh

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
		Mongosh: Mongosh{
			MongoHost:       viper.GetString("mongosh.host"),
			MongoPort:       viper.GetInt("mongosh.port"),
			MongoDatabase:   viper.GetString("mongosh.database"),
			MongoCollection: viper.GetString("mongosh.collection"),
		},
		MessageServiceHost: viper.GetString("server.host"),
		MessageServicePort: viper.GetInt("server.port"),
	}
	return &conf, nil
}
