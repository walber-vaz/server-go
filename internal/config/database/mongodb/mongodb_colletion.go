package mongodb

import (
	"context"
	"server-go/internal/config"
	"server-go/internal/config/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitConnection(config *config.DB) {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI(config.Host + ":" + config.Port))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	logger.Info("[INFO] Connected to MongoDB")
}
