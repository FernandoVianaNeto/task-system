package repository

import (
	"context"
	"task-system/internal/domain/entities"
	"task-system/internal/infrastructure/repository/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, userUuid string, input entities.Task) error {
	result := r.db.WithContext(ctx).Create(models.Task{
		Id:       input.Id,
		Uuid:     input.Uuid,
		UserUuid: input.User.Uuid,
		Title:    input.Title,
		Summary:  input.Summary,
		Status:   input.Status,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (r *TaskRepository) GetTaskByUser(ctx context.Context, taskUuid string, userUuid string) (*entities.Task, error) {
// 	result := r.db.WithContext(ctx).Create(&input)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
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
