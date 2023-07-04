package repositories

type IUserRepo interface {
	CreateUser(ctx context.Context, User *entities.User) error
	UpdateUser(ctx context.Context, updateFields []string, User *entities.User) error
	GetUser(ctx context.Context, id entity.ID) (*entities.User, error)
	GetUsers(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteUser(ctx context.Context, id entity.ID) error
	Log(ctx context.Context) *zap.SugaredLogger
}

type UserRepo struct {
}

func (u *UserRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.User{})
}

func (u UserRepo) CreateUser(ctx context.Context, user *entities.User) error {
	panic("todo")
}

func (u UserRepo) UpdateUser(ctx context.Context, updateFields []string, user *entities.User) error {
	panic("todo")
}

func (u UserRepo) GetUser(ctx context.Context, id entity.ID) (*entities.User, error) {
	panic("todo")
}

func (u UserRepo) DeleteUser(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (u UserRepo) GetUsers(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error) {
	panic("todo")
}

func (u UserRepo) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("UserRepo")
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}
