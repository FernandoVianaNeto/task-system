package domain_usecase

type Usecases struct {
	CreateTaskUsecase       CreateTaskUseCaseInterface
	CreateUserUsecase       CreateUserUsecaseInterface
	GetUserUsecase          GetUserUsecaseInterface
	AuthUsecase             AuthUsecaseInterface
	ListTaskUsecase         ListTaskUsecaseInterface
	UpdateTaskStatusUsecase UpdateTaskStatusUsecaseInterface
	DeleteTaskUsecase       DeleteTaskUsecaseInterface
}
