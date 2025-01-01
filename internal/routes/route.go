package routes

import (
	"connectrpc.com/grpcreflect"
	"github.com/shooters/user/internal/gen/protos/shooters/user/v1/pbconnect"
)

type routeNamer struct {
	names []string
}

func (n *routeNamer) Names() []string {
	return n.names
}

func (n *routeNamer) Set(names ...string) {
	for _, name := range names {
		n.names = append(n.names, name)
	}
}

func NewNamer() grpcreflect.Namer {
	return &routeNamer{
		names: PublicRoutes,
	}
}

// PublicRoutes Define public methods that don't require authentication
var PublicRoutes = []string{
	"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo",
	grpcreflect.ReflectV1ServiceName,
	grpcreflect.ReflectV1AlphaServiceName,
	pbconnect.UserServiceCreateUserProcedure,
	pbconnect.UserServiceListUsersProcedure,
	pbconnect.UserServiceGetUserProcedure,
}
