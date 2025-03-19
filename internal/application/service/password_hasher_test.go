package service_test

import (
	"task-system/internal/application/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHasherService_HashPassword(t *testing.T) {
	passwordHasher := service.NewPasswordHasherService()

	t.Run("should successfully hash a password", func(t *testing.T) {
		password := "securepassword"
		hashedPassword, err := passwordHasher.HashPassword(password, bcrypt.DefaultCost)

		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)
		assert.NotEqual(t, password, string(hashedPassword))

		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
		assert.NoError(t, err)
	})

	t.Run("should return an error if the bcrypt cost is too high", func(t *testing.T) {
		password := "securepassword"
		_, err := passwordHasher.HashPassword(password, 100)

		assert.Error(t, err)
	})
}
