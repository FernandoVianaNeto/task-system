package repository_user

import (
	"context"
	"task-system/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, input entities.User) error {
	user := User{
		Uuid: input.Uuid,
		Role: input.Role,
		Name: input.Name,
	}

	result := r.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (r *TaskRepository) GetTaskByUser(ctx context.Context, taskUuid string, userUuid string) (*entities.Task, bool) {
// 	result, err := r.db.WithContext(ctx).Get(taskUuid)

// 	if err == false {
// 		return &entities.Task{}, err
// 	}

// 	return result, err
// }

// func (r *TaskRepository) UpdateTaskByUser(ctx context.Context, userUuid string, input entities.Task) error {
// 	result := r.db.WithContext(ctx).
// 		Model(&models.PlanExtension{}).
// 		Where("freight_id = ? AND disabled = ?", freightId, false).
// 		Update("disabled", true).
// 		Update("disable_reason", disableReason)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	log.Println(fmt.Sprintf("Extensão desativada para o frete %d", freightId))

// 	return nil
// }
