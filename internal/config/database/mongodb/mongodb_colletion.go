package mongodb

import (
	"context"
	"server-go/internal/config"
	"server-go/internal/config/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDBConnection(ctx context.Context, config *config.DB) (*mongo.Database, error) {
	logger.Info("[INFO] Conectando com o MongoDB")

	client, err := mongo.Connect(options.Client().ApplyURI(config.Host + ":" + config.Port))
	if err != nil {
		logger.Error("[ERROR] Erro ao conectar com o MongoDB", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("[ERROR] Erro ao pingar o MongoDB", err)
		return nil, err
	}

	logger.Info("[INFO] Conexão com o MongoDB realizada com sucesso")
	return client.Database(config.Name), nil
}
