package repositories

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/shooters/user/internal/types"
	"gorm.io/gorm"
	"log"
	"os"
)

type keycloakRepository struct {
	engine       *gocloak.GoCloak
	store        *gorm.DB
	clientId     string
	clientSecret string
	realm        string
}

func (repository *keycloakRepository) CreateUser(ctx context.Context, user *gocloak.User) (interface{}, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to login: %v", err)))
	}

	userId, err := repository.engine.CreateUser(ctx, token, repository.realm, *user)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to create user: %v", err)))
	}

	return repository.GetUserById(ctx, userId)
}

func (repository *keycloakRepository) GetUsers(ctx context.Context) (interface{}, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to login: %v", err)))
	}
	users, err := repository.engine.GetUsers(ctx, token, repository.realm, gocloak.GetUsersParams{
		BriefRepresentation: gocloak.BoolP(false),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to get users: %v", err)))
	}

	return users, nil
}

func (repository *keycloakRepository) GetUserById(ctx context.Context, s string) (interface{}, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to login: %v", err)))
	}

	user, err := repository.engine.GetUserByID(ctx, token, repository.realm, s)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New(fmt.Sprintf("unable to get user: %v", err)))
	}

	log.Println("KcUser :::: |", user.Groups)

	return user, nil
}

func (repository *keycloakRepository) loginAdmin(ctx context.Context) (string, error) {
	jwtToken, err := repository.engine.LoginClient(ctx, repository.clientId, repository.clientSecret, repository.realm)
	if err != nil {
		return "", err
	}

	return jwtToken.AccessToken, nil
}

func NewKeycloakUserRepository(keycloak *gocloak.GoCloak, engine *gorm.DB) types.UserRepository {
	return &keycloakRepository{
		store:        engine,
		engine:       keycloak,
		clientId:     os.Getenv("KEYCLOAK.CLIENT_ID"),
		clientSecret: os.Getenv("KEYCLOAK.CLIENT_SECRET"),
		realm:        os.Getenv("KEYCLOAK.REALM"),
	}
}
