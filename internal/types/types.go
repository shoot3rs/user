package types

import (
	"connectrpc.com/connect"
	"context"
	"github.com/Nerzal/gocloak/v13"
	pb "github.com/shoot3rs/user/gen/shooters/user/v1"
	"github.com/shoot3rs/user/internal/auth"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
)

type ContextHelper interface {
	GetTenant(context.Context) (string, error)
	GetUserClaims(context.Context) *auth.UserAuthClaims
	GetUserClaimsFromRequest(*connect.Request[connect.AnyRequest]) *auth.UserAuthClaims
}

type GlobalConfig interface {
	GetGorm() *gorm.Config
	GetServerAddr() string
	GetServerConfig() *http2.Server
	LoadEnv()
	Logger() *zap.Logger
}

type Connection interface {
	Connect()
	GetConfig() *gorm.Config
	GetEngine() interface{}
}

type GRPCAuthMiddleware interface {
	CorsMiddleware(http.Handler) http.Handler
	LoggingUnaryInterceptor() grpc.UnaryServerInterceptor
	StreamInterceptor(auth.ServiceAuthenticator) grpc.StreamServerInterceptor
	TenantHeaderInterceptor([]string) grpc.UnaryServerInterceptor
	TokenValidationUnaryInterceptor(auth.ServiceAuthenticator) grpc.UnaryServerInterceptor
	UnaryTenantHeaderInterceptor(auth.ServiceAuthenticator) connect.UnaryFunc
	UnaryTokenInterceptor(auth.ServiceAuthenticator) connect.UnaryInterceptorFunc
}

type UserRepository interface {
	CreateUser(context.Context, *gocloak.User) (interface{}, error)
	GetUsers(context.Context, *connect.Request[pb.ListUsersRequest]) (interface{}, error)
	GetUserById(context.Context, string) (interface{}, error)
}

type UserService interface {
	CreateUser(context.Context, *connect.Request[pb.CreateUserRequest]) (*pb.User, error)
	GetUserById(context.Context, *connect.Request[pb.GetUserRequest]) (*pb.User, error)
	ListUsers(context.Context, *connect.Request[pb.ListUsersRequest]) ([]*pb.User, error)
}
