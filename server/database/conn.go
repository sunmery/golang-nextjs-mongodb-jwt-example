package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client  *mongo.Client
	ConnErr error
)

const URI = "mongodb://root:msdnmm@192.168.0.152:27017/"

func Start() {
	// db s
	Client, ConnErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	// db e
}
