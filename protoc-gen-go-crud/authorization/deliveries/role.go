package authorizationV1

import (
	components "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/api/v1/components"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type RoleService struct {
}

func (a RoleService) CreateRole(*gin.Context, *authorizationV1.CreateRoleRequest) (*empty.Empty, error) {
	panic("todo")
}

func (a RoleService) EnableRole(*gin.Context, *authorizationV1.EnableRoleRequest) (*empty.Empty, error) {
	panic("todo")
}

func (a RoleService) GetRole(*gin.Context, *authorizationV1.GetRoleRequest) (*components.Role, error) {
	panic("todo")
}

func (a RoleService) ListRoles(*gin.Context, *authorizationV1.ListRolesRequest) (*authorizationV1.Roles, error) {
	panic("todo")
}

func (a RoleService) UpdateRole(*gin.Context, *authorizationV1.UpdateRoleRequest) (*empty.Empty, error) {
	panic("todo")
}
