package repository

// import (
// 	"context"
// 	"errors"
// 	"log"

// )

// type TaskRepository struct {
// }

// func NewFreightPlanExtensionRepository(db *mongo.Database) domain_repository.PayPerUseRepositoryInterface {
// 	collection := db.Collection(configs.MongoCfg.PayPerUseCollection)

// 	return &FreightPlanExtensionRepository{
// 		db:         db,
// 		collection: collection,
// 	}
// }

// func (f *FreightPlanExtensionRepository) GetActivePlanExtensionByFreightId(ctx context.Context, freightId int) (*entity.PlanExtension, error) {
// 	var model FreightPlanExtensionModel

// 	filter := bson.M{
// 		"freight_id": freightId,
// 		"disabled":   false,
// 	}

// 	err := f.collection.FindOne(ctx, filter).Decode(&model)

// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, nil
// 		}

// 		return nil, err
// 	}

// 	entity := entity.PlanExtension{
// 		FreightID:           model.FreightID,
// 		Type:                model.Type,
// 		CreatedAt:           model.CreatedAt,
// 		CompanyId:           model.CompanyId,
// 		UserId:              model.UserId,
// 		TruckersExtraAmount: model.TruckersExtraAmount,
// 		Disabled:            model.Disabled,
// 	}

// 	return &entity, nil
// }

// func (f *FreightPlanExtensionRepository) CreatePlanExtension(ctx context.Context, input entity.PlanExtension) error {
// 	_, err := f.collection.InsertOne(ctx, FreightPlanExtensionModel{
// 		FreightID:           input.FreightID,
// 		CreatedAt:           input.CreatedAt,
// 		Type:                input.Type,
// 		CompanyId:           input.CompanyId,
// 		TruckersExtraAmount: input.TruckersExtraAmount,
// 		UserId:              input.UserId,
// 		FreeConsumption:     input.FreeConsumption,
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (f *FreightPlanExtensionRepository) DisableExtension(ctx context.Context, input entity.DisableExtensionEntity) error {
// 	filter := bson.M{
// 		"freight_id": input.FreightId,
// 		"disabled":   false,
// 	}

// 	updateFreight := bson.M{
// 		"$set": bson.M{
// 			"disabled":       input.Disabled,
// 			"disable_reason": input.DisableReason,
// 		},
// 	}

// 	result, err := f.collection.UpdateOne(
// 		ctx,
// 		filter,
// 		updateFreight,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	log.Println(result.ModifiedCount)

// 	return nil
// }
