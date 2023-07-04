package deliveries

import (
	components "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/api/v1/components"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	uc usecase.IUserUseCase
}

func (u UserService) ChangePassword(c *gin.Context, in *usersV1.ChangePasswordRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) CreateUser(c *gin.Context, in *usersV1.CreateUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) DeleteUser(c *gin.Context, in *usersV1.DeleteUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) EnableUser(c *gin.Context, in *usersV1.EnableUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) GetUser(c *gin.Context, in *usersV1.GetUserRequest) (*components.User, error) {
	panic("todo")
}

func (u UserService) ListUsers(c *gin.Context, in *usersV1.ListUsersRequest) (*usersV1.ListUsersReply, error) {
	panic("todo")
}

func (u UserService) Reset2FA(c *gin.Context, in *usersV1.Reset2FARequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) ResetActivate(c *gin.Context, in *usersV1.ResetActivateRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) ResetPassword(c *gin.Context, in *usersV1.ResetPasswordRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) UpdateCurrentUser(c *gin.Context, in *usersV1.UpdateCurrentUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) UpdateUser(c *gin.Context, in *usersV1.UpdateUserRequest) (*empty.Empty, error) {
	panic("todo")
}

func (u UserService) UploadProfile(c *gin.Context, in *usersV1.UploadProfileRequest) (*usersV1.UploadProfileResponse, error) {
	panic("todo")
}

func (u UserService) User(c *gin.Context, in *empty.Empty) (*components.User, error) {
	panic("todo")
}

func (u UserService) Log(c *gin.Context) *zap.SugaredLogger {
	return global.Logger(c).Named("UserRepo")
}

func NewUserService(uc usecase.IUserUseCase) usersV1.UserHTTPServer {
	return &UserService{
		uc: uc,
	}
}
