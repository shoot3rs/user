package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/Nerzal/gocloak/v13"
	"github.com/shoot3rs/user/internal/auth"
	"github.com/shoot3rs/user/internal/config"
	"github.com/shoot3rs/user/internal/data/repositories"
	"github.com/shoot3rs/user/internal/database"
	"github.com/shoot3rs/user/internal/gen/protos/shooters/user/v1/userv1connect"
	"github.com/shoot3rs/user/internal/helpers"
	"github.com/shoot3rs/user/internal/middlewares"
	"github.com/shoot3rs/user/internal/servers"
	"github.com/shoot3rs/user/internal/services"
	"golang.org/x/net/http2/h2c"
	"gorm.io/gorm"
	"net/http"

	"log"
	"os"
)

var (
	cfg = config.New()
	db  = database.New(cfg.GetGorm())
)

func init() {
	cfg.LoadEnv()
	db.Connect()
}

func main() {
	serverAddr := cfg.GetServerAddr()

	dbEngine := db.GetEngine().(*gorm.DB)
	authMiddleware := middlewares.New()
	authenticator, err := auth.New()
	if err != nil {
		log.Println("Error creating authenticator:", err)
		os.Exit(1)
	}

	goCloak := gocloak.NewClient(os.Getenv("KEYCLOAK.URL"))
	contextHelper := helpers.NewContextHelper(authenticator)
	userRepository := repositories.NewKeycloakUserRepository(goCloak, dbEngine)
	userService := services.NewUserService(userRepository, contextHelper)
	userServer := servers.NewUserServer(userService)

	mux := http.NewServeMux()
	tokenInterceptor := authMiddleware.UnaryTokenInterceptor(authenticator)

	reflector := grpcreflect.NewStaticReflector(userv1connect.UserServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	interceptors := connect.WithInterceptors(tokenInterceptor)
	path, handler := userv1connect.NewUserServiceHandler(userServer, interceptors)
	handler = authMiddleware.CorsMiddleware(handler)
	mux.Handle(path, handler)

	log.Println("gRPC server started on port:", serverAddr)
	if err := http.ListenAndServe(
		serverAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, cfg.GetServerConfig()),
	); err != nil {
		log.Println("failed to start server:", err)
		os.Exit(1)
	}
}
