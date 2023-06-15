package repositories

type IRoleRepo interface {
	CreateRole(ctx context.Context, Role *entities.Role) error
	UpdateRole(ctx context.Context, updateFields []string, Role *entities.Role) error
	GetRole(ctx context.Context, id entity.ID) (*entities.Role, error)
	GetRoles(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteRole(ctx context.Context, id entity.ID) error
}

type RoleRepo struct {
}

func (m *RoleRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.Role{})
}

func (m RoleRepo) CreateRole(ctx context.Context, Role *entities.Role) error {
	panic("todo")
}

func (m RoleRepo) UpdateRole(ctx context.Context, updateFields []string, Role *entities.Role) error {
	panic("todo")
}

func (m RoleRepo) GetRole(ctx context.Context, id entity.ID) (*entities.Role, error) {
	panic("todo")
}

func (m RoleRepo) DeleteRole(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (m RoleRepo) GetRoles(ctx context.Context, filter *database.Filter) (res []*entities.Role, count int64, err error) {
	panic("todo")
}
