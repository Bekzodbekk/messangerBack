package mongosh

import (
	"context"
	"fmt"
	config "user-service/internal/user/pkg/load"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoshConn struct {
	Cli  *mongo.Client
	Coll *mongo.Collection
}

func InitDB(conf *config.Config) (*MongoshConn, error) {
	ctx := context.Background()

	uri := fmt.Sprintf("mongodb://%s:%d", conf.Mongosh.MongoHost, conf.Mongosh.MongoPort)
	// uri := fmt.Sprintf("mongodb+srv://%s:%s@fornt.otm6nho.mongodb.net/?retryWrites=true&w=majority&appName=Fornt",
	// 	conf.Mongosh.MongoUser,
	// 	conf.Mongosh.MongoPassword,
	// )
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx, nil); err != nil {
		return nil, err
	}

	coll := cli.Database(conf.Mongosh.MongoDatabase).Collection(conf.Mongosh.MongoCollection)

	return &MongoshConn{
		Cli:  cli,
		Coll: coll,
	}, nil
}
