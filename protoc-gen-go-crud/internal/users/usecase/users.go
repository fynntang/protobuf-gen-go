package usecase

type IUserUseCase interface {
	CreateUser(ctx context.Context, in *usersV1.CreateUserRequest) error
	UpdateUser(ctx context.Context, in *usersV1.UpdateUserRequest) error
	DeleteUser(ctx context.Context, in *usersV1.DeleteUserRequest) error
	GetUser(ctx context.Context, in *usersV1.GetUserRequest) (*entities.User, error)
	ListUser(ctx context.Context, in *usersV1.ListUserRequest) ([]*entities.User, int64, error)
	Log(ctx context.Context) *zap.SugaredLogger
}

type UserUseCase struct {
}

func NewUserUseCase() IUserUseCase {
	return &UserUseCase{}
}

func (u *UserUseCase) CreateUser(ctx context.Context, in *usersV1.CreateUserRequest) error {
	panic("todo")
}

func (u *UserUseCase) UpdateUser(ctx context.Context, in *usersV1.UpdateUserRequest) error {
	panic("todo")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, in *usersV1.DeleteUserRequest) error {
	panic("todo")
}

func (u *UserUseCase) GetUser(ctx context.Context, in *usersV1.GetUserRequest) (*entities.User, error) {
	panic("todo")
}

func (u *UserUseCase) ListUser(ctx context.Context, in *usersV1.ListUserRequest) ([]*entities.User, int64, error) {
	panic("todo")
}

func (u UserUseCase) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("UserRepo")
}
