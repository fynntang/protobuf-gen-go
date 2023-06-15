package deliveries

import (
	empty "github.com/golang/protobuf/ptypes/empty"
)

type AuthService struct {
}

func (a AuthService) ListMenus(*gin.Context, *empty.Empty) (*authorizationV1.Menus, error) {
	panic("todo")
}
