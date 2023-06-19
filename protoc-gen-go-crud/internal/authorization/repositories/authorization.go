package repositories

type IAuthRepo interface {
	CreateAuth(ctx context.Context, Auth *entities.Auth) error
	UpdateAuth(ctx context.Context, updateFields []string, Auth *entities.Auth) error
	GetAuth(ctx context.Context, id entity.ID) (*entities.Auth, error)
	GetAuths(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteAuth(ctx context.Context, id entity.ID) error
	Log(ctx context.Context) *zap.SugaredLogger
}

type AuthRepo struct {
}

func (a *AuthRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.Auth{})
}

func (a AuthRepo) CreateAuth(ctx context.Context, auth *entities.Auth) error {
	panic("todo")
}

func (a AuthRepo) UpdateAuth(ctx context.Context, updateFields []string, auth *entities.Auth) error {
	panic("todo")
}

func (a AuthRepo) GetAuth(ctx context.Context, id entity.ID) (*entities.Auth, error) {
	panic("todo")
}

func (a AuthRepo) DeleteAuth(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (a AuthRepo) GetAuths(ctx context.Context, filter *database.Filter) (res []*entities.Auth, count int64, err error) {
	panic("todo")
}

func (a AuthRepo) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("AuthRepo")
}
