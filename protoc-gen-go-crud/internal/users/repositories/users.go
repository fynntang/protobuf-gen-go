package repositories

type IUsersRepo interface {
	CreateUsers(ctx context.Context, Users *entities.Users) error
	UpdateUsers(ctx context.Context, updateFields []string, Users *entities.Users) error
	GetUsers(ctx context.Context, id entity.ID) (*entities.Users, error)
	GetUserss(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteUsers(ctx context.Context, id entity.ID) error
	Log(ctx context.Context) *zap.SugaredLogger
}

type UsersRepo struct {
}

func (u *UsersRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.Users{})
}

func (u UsersRepo) CreateUsers(ctx context.Context, users *entities.Users) error {
	panic("todo")
}

func (u UsersRepo) UpdateUsers(ctx context.Context, updateFields []string, users *entities.Users) error {
	panic("todo")
}

func (u UsersRepo) GetUsers(ctx context.Context, id entity.ID) (*entities.Users, error) {
	panic("todo")
}

func (u UsersRepo) DeleteUsers(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (u UsersRepo) GetUserss(ctx context.Context, filter *database.Filter) (res []*entities.Users, count int64, err error) {
	panic("todo")
}

func (u UsersRepo) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("UsersRepo")
}
