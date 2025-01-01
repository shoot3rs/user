package servers

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/shooters/user/internal/gen/protos/shooters/user/v1"
	"github.com/shooters/user/internal/gen/protos/shooters/user/v1/pbconnect"
	"github.com/shooters/user/internal/types"
)

type userServer struct {
	userService types.UserService
}

func (server *userServer) ListUsers(ctx context.Context, c *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	users, err := server.userService.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.ListUsersResponse]{
		Msg: &v1.ListUsersResponse{Users: users},
	}, nil
}

func (server *userServer) CreateUser(ctx context.Context, c *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	user, err := server.userService.CreateUser(ctx, c)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.CreateUserResponse]{
		Msg: &v1.CreateUserResponse{User: user},
	}, nil
}

func (server *userServer) GetUser(ctx context.Context, c *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	user, err := server.userService.GetUserById(ctx, c)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.GetUserResponse]{
		Msg: &v1.GetUserResponse{User: user},
	}, nil
}

func NewUserServer(userService types.UserService) pbconnect.UserServiceHandler {
	return &userServer{
		userService: userService,
	}
}
