package config

type MongoConfig struct {
	Host string
	Port int64
	Database, User, Pwd string
	AuthType string
}
