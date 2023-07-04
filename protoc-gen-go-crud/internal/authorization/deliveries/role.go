package deliveries

import (
	components "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/api/v1/components"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type RoleService struct {
	rc usecase.IRoleUseCase
}

func (r RoleService) CreateRole(c *gin.Context, in *authorizationV1.CreateRoleRequest) (*empty.Empty, error) {
	panic("todo")
}

func (r RoleService) EnableRole(c *gin.Context, in *authorizationV1.EnableRoleRequest) (*empty.Empty, error) {
	panic("todo")
}

func (r RoleService) GetRole(c *gin.Context, in *authorizationV1.GetRoleRequest) (*components.Role, error) {
	panic("todo")
}

func (r RoleService) ListRoles(c *gin.Context, in *authorizationV1.ListRolesRequest) (*authorizationV1.Roles, error) {
	panic("todo")
}

func (r RoleService) UpdateRole(c *gin.Context, in *authorizationV1.UpdateRoleRequest) (*empty.Empty, error) {
	panic("todo")
}

func (r RoleService) Log(c *gin.Context) *zap.SugaredLogger {
	return global.Logger(c).Named("RoleRepo")
}

func NewRoleService(rc usecase.IRoleUseCase) authorizationV1.RoleHTTPServer {
	return &RoleService{
		rc: rc,
	}
}
