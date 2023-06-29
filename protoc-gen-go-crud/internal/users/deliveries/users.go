package deliveries

import (
	components "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/api/v1/components"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type UsersService struct {
}

func (u UsersService) ChangePassword(c *gin.Context, in *usersV1.ChangePasswordRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) CreateUser(c *gin.Context, in *usersV1.CreateUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) DeleteUser(c *gin.Context, in *usersV1.DeleteUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) EnableUser(c *gin.Context, in *usersV1.EnableUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) GetUser(c *gin.Context, in *usersV1.GetUserRequest) (*components.User, error) {
	panic("todo")
}

func (u UsersService) ListUsers(c *gin.Context, in *usersV1.ListUsersRequest) (*usersV1.ListUsersReply, error) {
	panic("todo")
}

func (u UsersService) Reset2FA(c *gin.Context, in *usersV1.Reset2FARequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) ResetActivate(c *gin.Context, in *usersV1.ResetActivateRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) ResetPassword(c *gin.Context, in *usersV1.ResetPasswordRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) UpdateCurrentUser(c *gin.Context, in *usersV1.UpdateCurrentUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) UpdateUser(c *gin.Context, in *usersV1.UpdateUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UsersService) UploadProfile(c *gin.Context, in *usersV1.UploadProfileRequest) (*usersV1.UploadProfileResponse, error) {
	panic("todo")
}

func (u UsersService) User(c *gin.Context, in *empty.Empty) (*components.User, error) {
	panic("todo")
}

func (u UsersService) Log(c *gin.Context) *zap.SugaredLogger {
	return global.Logger(c).Named("UsersRepo")
}
