package configs

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGODB_URI")
}

func EnvMongoDatabaseName() string { return os.Getenv("MONGODB_NAME") }

func EnvJWTSecret() string { return os.Getenv("JWT_SECRET") }
