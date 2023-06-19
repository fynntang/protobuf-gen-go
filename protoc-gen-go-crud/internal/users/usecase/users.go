package usecase

type IUsersUsecase interface {
	CreateUsers(ctx context.Context, in *usersV1.CreateUsersRequest) error
	UpdateUsers(ctx context.Context, in *usersV1.UpdateUsersRequest) error
	DeleteUsers(ctx context.Context, in *usersV1.DeleteUsersRequest) error
	GetUsers(ctx context.Context, in *usersV1.GetUsersRequest) (*entities.Users, error)
	ListUsers(ctx context.Context, in *usersV1.ListUsersRequest) ([]*entities.Users, int64, error)
	Log(ctx context.Context) *zap.SugaredLogger
}

type UsersUseCase struct {
}

func NewUsersUseCase() IUsersUsecase {
	return &UsersUseCase{}
}

func (u *UsersUseCase) CreateUsers(ctx context.Context, in *usersV1.CreateUsersRequest) error {
	panic("todo")
}

func (u *UsersUseCase) UpdateUsers(ctx context.Context, in *usersV1.UpdateUsersRequest) error {
	panic("todo")
}

func (u *UsersUseCase) DeleteUsers(ctx context.Context, in *usersV1.DeleteUsersRequest) error {
	panic("todo")
}

func (u *UsersUseCase) GetUsers(ctx context.Context, in *usersV1.GetUsersRequest) (*entities.Users, error) {
	panic("todo")
}

func (u *UsersUseCase) ListUsers(ctx context.Context, in *usersV1.ListUsersRequest) ([]*entities.Users, int64, error) {
	panic("todo")
}

func (u UsersUseCase) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("UsersRepo")
}
