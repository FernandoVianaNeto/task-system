package app

import (
	"context"
	"task-system/internal/infrastructure/web"
)

type Application struct {
	UseCases UseCases
}

type Clients struct {
	planApi    client.CPMApiClientInterface
	freightApi client.FreightApiClientInterface
}

type UseCases struct {
}

func NewApplication() *web.Server {
	ctx := context.Background()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	repository := mongo_repository.NewFreightPlanExtensionRepository(db)

	clients := NewClients()

	usecases := NewUseCases(ctx, clients, repository)

	hasher := hasher.InitializeHasher()
	hash := hash.NewHasherAdapter(hasher)

	srv := web.NewServer(ctx)

	return srv
}
