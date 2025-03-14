package repository_task

import (
	"context"
	"task-system/internal/domain/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, input entities.Task) error {
	task := Task{
		Id:       input.Id,
		Uuid:     input.Uuid,
		UserUuid: input.UserUuid,
		Title:    input.Title,
		Summary:  input.Summary,
		Status:   input.Status,
	}

	result := r.db.WithContext(ctx).Create(&task)

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
