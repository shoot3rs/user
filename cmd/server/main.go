package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/Nerzal/gocloak/v13"
	"github.com/shooters/user/internal/auth"
	"github.com/shooters/user/internal/config"
	"github.com/shooters/user/internal/data/repositories"
	"github.com/shooters/user/internal/database"
	"github.com/shooters/user/internal/gen/protos/shooters/user/v1/pbconnect"
	"github.com/shooters/user/internal/helpers"
	"github.com/shooters/user/internal/middlewares"
	"github.com/shooters/user/internal/routes"
	"github.com/shooters/user/internal/servers"
	"github.com/shooters/user/internal/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/gorm"
	"net/http"

	"log"
	"os"
)

var (
	cfg = config.New()
	db  = database.NewDatabase(cfg.GetGorm())
)

func init() {
	cfg.LoadEnv()
	db.Connect()
}

func main() {
	serverAddr := cfg.GetServerAddr()
	authMiddleware := middlewares.New()
	authenticator, err := auth.New()
	if err != nil {
		log.Println("Error creating authenticator:", err)
		os.Exit(1)
	}

	dbEngine := db.GetEngine().(*gorm.DB)

	goCloak := gocloak.NewClient(os.Getenv("KEYCLOAK.URL"))
	ctxHelper := helpers.NewGrpcRequestHelper(authenticator)
	userRepository := repositories.NewKeycloakUserRepository(goCloak, dbEngine)
	categoryService := services.NewUserService(userRepository, ctxHelper)
	categoryServer := servers.NewUserServer(categoryService)

	mux := http.NewServeMux()
	tokenInterceptor := authMiddleware.UnaryTokenInterceptor(authenticator)

	serviceNames := routes.NewNamer()
	reflector := grpcreflect.NewStaticReflector(serviceNames.Names()...)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	interceptors := connect.WithInterceptors(tokenInterceptor)
	path, handler := pbconnect.NewUserServiceHandler(categoryServer, interceptors)
	handler = authMiddleware.CorsMiddleware(handler)
	mux.Handle(path, handler)

	log.Println("gRPC server started on port:", serverAddr)
	if err := http.ListenAndServe(
		serverAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Println("failed to start server:", err)
		os.Exit(1)
	}
}
