package domain_service

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type PasswordHasherServiceInterface interface {
	HashPassword(password string, cost int) ([]byte, error)
}
