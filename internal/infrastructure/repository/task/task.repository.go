package repository_task

import (
	"context"
	"fmt"
	"task-system/internal/domain/dto"
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
		Uuid:    input.Uuid,
		Owner:   input.Owner,
		Title:   input.Title,
		Summary: input.Summary,
		Status:  input.Status,
	}

	result := r.db.WithContext(ctx).Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) ListTask(ctx context.Context, input dto.ListTaskDto) ([]*entities.Task, error) {
	var entity []*entities.Task
	var result *gorm.DB

	query := r.db.WithContext(ctx)

	fmt.Println("INPUT", input)

	if input.Role == "developer" {
		if input.Uuid != nil {
			query = query.Where("uuid = ? AND owner = ?", input.Uuid, input.UserUuid)
		}
		fmt.Println("INPUT", input)

		query = query.Where("owner = ?", input.UserUuid)
	}

	if input.Role == "admin" {
		if input.Uuid != nil {
			query = query.Where("uuid = ?", input.Uuid)
		}

		if input.Owner != nil {
			query = query.Where("owner = ?", input.Owner)
		}
	}

	result = query.Find(&entity)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity, nil
		}
		return nil, result.Error
	}

	return entity, result.Error
}

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
