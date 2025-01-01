package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	v1 "github.com/shooters/user/internal/gen/protos/shooters/user/v1"
	"github.com/shooters/user/internal/types"
	"github.com/shooters/user/utils"
	"log"
)

type userService struct {
	ctxHelper      types.RequestHelper
	userRepository types.UserRepository
}

func (service *userService) CreateUser(ctx context.Context, request *connect.Request[v1.CreateUserRequest]) (*v1.User, error) {
	newUser, err := utils.NewUserFromRequest(request.Msg.GetUser())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("error creating new user: %v", err)))
	}

	log.Println("Creating new user :::: |", newUser)
	user, err := service.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("failed to create user: %v", err)))
	}

	kcUser := user.(*gocloak.User)

	userProto, err := utils.NewProtoFromKCUser(kcUser)
	if err != nil {
		return nil, err
	}

	return userProto, nil
}

func (service *userService) GetUserById(ctx context.Context, request *connect.Request[v1.GetUserRequest]) (*v1.User, error) {
	user, err := service.userRepository.GetUserById(ctx, request.Msg.GetId())
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New(fmt.Sprintf("user not found: %v", err)))
	}

	kcUser := user.(*gocloak.User)

	userProto, err := utils.NewProtoFromKCUser(kcUser)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("failed to get user: %v", err)))
	}

	return userProto, nil
}

func (service *userService) GetAllUsers(ctx context.Context) ([]*v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository types.UserRepository, requestHelper types.RequestHelper) types.UserService {
	return &userService{
		ctxHelper:      requestHelper,
		userRepository: userRepository,
	}
}
