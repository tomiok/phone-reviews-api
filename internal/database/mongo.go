package database

import (
	"context"
	"fmt"
	"github.com/tomiok/course-phones-review/internal/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoClient(host string) *Mongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)))

	if err != nil {
		logs.Log().Errorf("cannot connect with mongo %s", err.Error())
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		logs.Log().Errorf("cannot connect with mongo %s", err.Error())
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		logs.Log().Errorf("cannot PING the mongo server: %s", err.Error())
		panic(err)
	}

	return &Mongo{client}
}
