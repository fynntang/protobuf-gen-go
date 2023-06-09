// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v0.0.1
// - protoc             v3.21.12
// source: v1/users/users.proto

package usersV1

import (
	components "bitbucket.org/antalphadev/earn/api/v1/components"
	bytes "bytes"
	json "encoding/json"
	gin "github.com/gin-gonic/gin"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
var _ = new(gin.Engine)
var _ = new(http.Request)
var _ = new(json.Encoder)
var _ = new(jsonpb.Marshaler)
var _ = new(proto.Message)
var _ = new(bytes.Buffer)

const RouteUsersChangePassword = "/v1/user/password/change"
const RouteUsersCreateUser = "/v1/users"
const RouteUsersDeleteUser = "/v1/users/:id"
const RouteUsersEnableUser = "/v1/users/:id/enable"
const RouteUsersGetUser = "/v1/users/:id"
const RouteUsersListUsers = "/v1/users"
const RouteUsersReset2FA = "/v1/users/:id/2fa/reset"
const RouteUsersResetActivate = "/v1/users/:id/activate"
const RouteUsersResetPassword = "/v1/users/:id/password/reset"
const RouteUsersUpdateCurrentUser = "/v1/user"
const RouteUsersUpdateUser = "/v1/users/:id"
const RouteUsersUploadProfile = "/v1/user/upload-profile"
const RouteUsersUser = "/v1/user"

type UsersHTTPServer interface {
	// ChangePassword 当前用户修改密码
	ChangePassword(*gin.Context, *ChangePasswordRequest) (*empty.Empty, error)
	// CreateUser 新增用户
	CreateUser(*gin.Context, *CreateUserRequest) (*empty.Empty, error)
	// DeleteUser 删除用户
	DeleteUser(*gin.Context, *DeleteUserRequest) (*empty.Empty, error)
	// EnableUser 用户启用/禁用
	EnableUser(*gin.Context, *EnableUserRequest) (*empty.Empty, error)
	// GetUser 查看用户详情
	GetUser(*gin.Context, *GetUserRequest) (*components.User, error)
	// ListUsers 用户列表
	ListUsers(*gin.Context, *ListUsersRequest) (*ListUsersReply, error)
	// Reset2FA 重置2fa
	Reset2FA(*gin.Context, *Reset2FARequest) (*empty.Empty, error)
	// ResetActivate 重新激活
	ResetActivate(*gin.Context, *ResetActivateRequest) (*empty.Empty, error)
	// ResetPassword 重置密码
	ResetPassword(*gin.Context, *ResetPasswordRequest) (*empty.Empty, error)
	// UpdateCurrentUser 当前用户更新
	UpdateCurrentUser(*gin.Context, *UpdateCurrentUserRequest) (*empty.Empty, error)
	// UpdateUser 更新用户
	UpdateUser(*gin.Context, *UpdateUserRequest) (*empty.Empty, error)
	// UploadProfile 上传头像
	UploadProfile(*gin.Context, *UploadProfileRequest) (*UploadProfileResponse, error)
	// User 当前用户
	User(*gin.Context, *empty.Empty) (*components.User, error)
}

func RegisterUsersHTTPServer(r gin.IRouter, srv UsersHTTPServer) {
	r.GET(RouteUsersListUsers, _Users_ListUsers0_HTTP_Handler(srv))
	r.POST(RouteUsersCreateUser, _Users_CreateUser0_HTTP_Handler(srv))
	r.PUT(RouteUsersUpdateUser, _Users_UpdateUser0_HTTP_Handler(srv))
	r.GET(RouteUsersGetUser, _Users_GetUser0_HTTP_Handler(srv))
	r.DELETE(RouteUsersDeleteUser, _Users_DeleteUser0_HTTP_Handler(srv))
	r.POST(RouteUsersUploadProfile, _Users_UploadProfile0_HTTP_Handler(srv))
	r.POST(RouteUsersResetActivate, _Users_ResetActivate0_HTTP_Handler(srv))
	r.POST(RouteUsersResetPassword, _Users_ResetPassword0_HTTP_Handler(srv))
	r.POST(RouteUsersReset2FA, _Users_Reset2FA0_HTTP_Handler(srv))
	r.POST(RouteUsersEnableUser, _Users_EnableUser0_HTTP_Handler(srv))
	r.GET(RouteUsersUser, _Users_User0_HTTP_Handler(srv))
	r.PUT(RouteUsersChangePassword, _Users_ChangePassword0_HTTP_Handler(srv))
	r.PUT(RouteUsersUpdateCurrentUser, _Users_UpdateCurrentUser0_HTTP_Handler(srv))
}

func _Users_ListUsers0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ListUsersRequest
		if err := c.BindQuery(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.ListUsers(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_CreateUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in CreateUserRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.CreateUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_UpdateUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in UpdateUserRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.UpdateUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_GetUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in GetUserRequest
		if err := c.BindQuery(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.GetUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_DeleteUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in DeleteUserRequest
		if err := c.BindQuery(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.DeleteUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_UploadProfile0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in UploadProfileRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.UploadProfile(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_ResetActivate0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ResetActivateRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.ResetActivate(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_ResetPassword0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ResetPasswordRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.ResetPassword(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_Reset2FA0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in Reset2FARequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.Reset2FA(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_EnableUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in EnableUserRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.EnableUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_User0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in empty.Empty
		if err := c.BindQuery(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.User(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_ChangePassword0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ChangePasswordRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.ChangePassword(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}

func _Users_UpdateCurrentUser0_HTTP_Handler(srv UsersHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in UpdateCurrentUserRequest

		if err := c.Bind(&in); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		out, err := srv.UpdateCurrentUser(c, &in)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData)
	}
}
