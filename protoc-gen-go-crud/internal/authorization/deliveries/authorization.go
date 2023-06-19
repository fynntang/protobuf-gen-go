package deliveries

import (
	empty "github.com/golang/protobuf/ptypes/empty"
)

type AuthService struct {
}

func (a AuthService) ListMenus(c *gin.Context, in *empty.Empty) (*authorizationV1.Menus, error) {
	panic("todo")
}

func (a AuthService) Log(c *gin.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("AuthRepo")
}
