package repositories

type IRoleRepo interface {
	CreateRole(ctx context.Context, Role *entities.Role) error
	UpdateRole(ctx context.Context, updateFields []string, Role *entities.Role) error
	GetRole(ctx context.Context, id entity.ID) (*entities.Role, error)
	GetRoles(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteRole(ctx context.Context, id entity.ID) error
	Log(ctx context.Context) *zap.SugaredLogger
}

type RoleRepo struct {
}

func (r *RoleRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.Role{})
}

func (r RoleRepo) CreateRole(ctx context.Context, role *entities.Role) error {
	panic("todo")
}

func (r RoleRepo) UpdateRole(ctx context.Context, updateFields []string, role *entities.Role) error {
	panic("todo")
}

func (r RoleRepo) GetRole(ctx context.Context, id entity.ID) (*entities.Role, error) {
	panic("todo")
}

func (r RoleRepo) DeleteRole(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (r RoleRepo) GetRoles(ctx context.Context, filter *database.Filter) (res []*entities.Role, count int64, err error) {
	panic("todo")
}

func (r RoleRepo) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("RoleRepo")
}
