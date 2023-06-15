package usecase

type IRoleUsecase interface {
	CreateRole(ctx context.Context, in *authorizationV1.CreateRoleRequest) error
	UpdateRole(ctx context.Context, in *authorizationV1.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, in *authorizationV1.DeleteRoleRequest) error
	GetRole(ctx context.Context, in *authorizationV1.GetRoleRequest) (*entities.Role, error)
	ListRole(ctx context.Context, in *authorizationV1.ListRoleRequest) ([]*entities.Role, int64, error)
}

type RoleUseCase struct {
}

func NewRoleUseCase() IRoleUsecase {
	return &RoleUseCase{}
}

func (r *RoleUseCase) CreateRole(ctx context.Context, in *authorizationV1.CreateRoleRequest) error {
	panic("todo")
}

func (r *RoleUseCase) UpdateRole(ctx context.Context, in *authorizationV1.UpdateRoleRequest) error {
	panic("todo")
}

func (r *RoleUseCase) DeleteRole(ctx context.Context, in *authorizationV1.DeleteRoleRequest) error {
	panic("todo")
}

func (r *RoleUseCase) GetRole(ctx context.Context, in *authorizationV1.GetRoleRequest) (*entities.Role, error) {
	panic("todo")
}

func (r *RoleUseCase) ListRole(ctx context.Context, in *authorizationV1.ListRoleRequest) ([]*entities.Role, int64, error) {
	panic("todo")
}
