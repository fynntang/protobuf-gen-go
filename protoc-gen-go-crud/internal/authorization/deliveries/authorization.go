package authorizationV1

import (
	authorizationV1 "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/api/v1/authorization"
	"github.com/gin-gonic/gin"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type AuthService struct {
}

func (a AuthService) ListMenus(*gin.Context, *empty.Empty) (*authorizationV1.Menus, error) {
	panic("todo")
}
