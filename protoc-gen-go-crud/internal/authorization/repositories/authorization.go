package repositories

import (
	"context"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/internal/authorization/entities"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/pkg/database"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/pkg/database/entity"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/pkg/global"
)

type IAuthRepo interface {
	CreateAuth(ctx context.Context, Auth *entities.Auth) error
	UpdateAuth(ctx context.Context, updateFields []string, Auth *entities.Auth) error
	GetAuth(ctx context.Context, id entity.ID) (*entities.Auth, error)
	GetAuths(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	DeleteAuth(ctx context.Context, id entity.ID) error
}

type AuthRepo struct {
}

func (m *AuthRepo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.Auth{})
}

func (m AuthRepo) CreateAuth(ctx context.Context, Auth *entities.Auth) error {
	panic("todo")
}

func (m AuthRepo) UpdateAuth(ctx context.Context, updateFields []string, Auth *entities.Auth) error {
	panic("todo")
}

func (m AuthRepo) GetAuth(ctx context.Context, id entity.ID) (*entities.Auth, error) {
	panic("todo")
}

func (m AuthRepo) DeleteAuth(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (m AuthRepo) GetAuths(ctx context.Context, filter *database.Filter) (res []*entities.Auth, count int64, err error) {
	panic("todo")
}
