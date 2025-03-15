package repository_user

import (
	"context"
	"fmt"
	"task-system/internal/domain/dto"
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
		Uuid:     input.Uuid,
		Role:     input.Role,
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}

	result := r.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) GetUserByUuid(ctx context.Context, input dto.GetUserByUuidDto) (*entities.User, error) {
	var user entities.User

	result := r.db.WithContext(ctx).Where("uuid = ?", input.Uuid).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &entities.User{}, nil
		}
		return nil, result.Error
	}

	return &user, result.Error
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, input dto.GetUserByEmailDto) (*entities.User, error) {
	var user entities.User

	fmt.Println("CHEGUEI AQUI", input)

	result := r.db.WithContext(ctx).Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &entities.User{}, nil
		}
		return nil, result.Error
	}

	return &user, result.Error
}
