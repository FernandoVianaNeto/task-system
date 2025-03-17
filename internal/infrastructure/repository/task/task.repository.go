package repository_task

import (
	"context"
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

	if input.Uuid != nil {
		result = r.db.WithContext(ctx).Where("uuid = ?", input.Uuid).Find(&entity)
	} else if input.Owner != nil {
		result = r.db.WithContext(ctx).Where("owner = ?", input.Owner).Find(&entity)
	} else {
		result = r.db.WithContext(ctx).Find(&entity)

	}

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
