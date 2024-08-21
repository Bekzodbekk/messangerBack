package load

import "github.com/spf13/viper"

type Mongosh struct {
	MongoHost       string
	MongoPort       int
	MongoDatabase   string
	MongoCollection string
}

type Redis struct {
	RedisHost string
	RedisPort int
}

type Config struct {
	Mongosh Mongosh

	Redis Redis

	UserServiceHost string
	UserServicePort int
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
		Redis: Redis{
			RedisHost: viper.GetString("redis.host"),
			RedisPort: viper.GetInt("redis.port"),
		},

		UserServiceHost: viper.GetString("server.host"),
		UserServicePort: viper.GetInt("server.port"),
	}
	return &conf, nil
}
