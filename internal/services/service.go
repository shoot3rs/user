package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	v1 "github.com/shoot3rs/user/gen/shooters/user/v1"
	"github.com/shoot3rs/user/internal/types"
	"github.com/shoot3rs/user/utils"
	"log"
)

type userService struct {
	ctxHelper      types.ContextHelper
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

func (service *userService) ListUsers(ctx context.Context, request *connect.Request[v1.ListUsersRequest]) ([]*v1.User, error) {
	userRepresentations, err := service.userRepository.GetUsers(ctx, request)
	if err != nil {
		return nil, err
	}

	userPbs := make([]*v1.User, 0)
	users := userRepresentations.([]*gocloak.User)
	for _, user := range users {
		kcUser, err := utils.NewProtoFromKCUser(user)
		if err != nil {
			log.Println("failed to convert kc user to proto:|", err)
		}

		userPbs = append(userPbs, kcUser)
	}

	return userPbs, nil
}

func NewUserService(userRepository types.UserRepository, requestHelper types.ContextHelper) types.UserService {
	return &userService{
		ctxHelper:      requestHelper,
		userRepository: userRepository,
	}
}
