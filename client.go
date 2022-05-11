package mongomgr

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(ctx context.Context) (*mongo.Database, error) {
	c, err := loadEnvConfig()
	if err != nil {
		return nil, err
	}
	uri := c.GetString("mongo.uri")
	if uri == "" {
		return nil, fmt.Errorf(
			"you must set your 'MONGO_URI' config",
		)
	}
	db := c.GetString("mongo.dbname")
	if db == "" {
		return nil, fmt.Errorf(
			"you must set your 'MONGO_DBNAME' config",
		)
	}
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	Client = conn
	return conn.Database(db), nil
}

func Disconnect(ctx context.Context) error {
	return Client.Disconnect(ctx)
}
