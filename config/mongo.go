package config

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

const CtxTimeOutDuration time.Duration = 30 * time.Second

type MongoConfig struct {
	Host string
	Port int64
	Database, User, Pwd string
	AuthType string
}

func NewMongoClient(uri string) (*mongo.Client, error) {
	return mongo.NewClient("mongodb://192.168.1.90:27017")
}
