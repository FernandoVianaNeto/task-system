package usecase

import (
	"context"
	"errors"
	configs "task-system/cmd/config"
	"task-system/internal/domain/dto"
	domain_repository "task-system/internal/domain/repository"
	domain_response "task-system/internal/domain/response"
	domain_usecase "task-system/internal/domain/usecase"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewAuthUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.AuthUsecaseInterface {
	return &AuthUsecase{
		UserRepository: repository,
	}
}

func (a *AuthUsecase) Execute(ctx context.Context, input dto.AuthDto) (domain_response.AuthResponse, error) {
	user, err := a.UserRepository.GetUserByEmail(ctx, dto.GetUserByEmailDto{Email: input.Email})

	if user.Uuid == "" {
		return domain_response.AuthResponse{}, errors.New("User not found")
	}

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	token, err := generateToken(input.Email, user.Role, user.Uuid)

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	return domain_response.AuthResponse{Token: token}, nil
}

func generateToken(email string, role string, userUuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_email": email,
		"user_role":  role,
		"user_uuid":  userUuid,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(configs.ApplicationCfg.JwtSecret))
}
