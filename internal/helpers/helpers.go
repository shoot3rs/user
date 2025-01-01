package helpers

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/shooters/user/internal/auth"
	"github.com/shooters/user/internal/types"
	"google.golang.org/grpc/metadata"
	"log"
)

type grpcRequestHelper struct {
	authenticator auth.ServiceAuthenticator
}

func (helper grpcRequestHelper) GetUserClaimsFromRequest(request *connect.Request[connect.AnyRequest]) *auth.UserAuthClaims {
	request.Header().Get("Authorization")

	return nil
}

func (helper grpcRequestHelper) GetUserClaims(ctx context.Context) *auth.UserAuthClaims {
	userClaims := ctx.Value(auth.ContextKeyUser).(*auth.UserAuthClaims)
	return userClaims
}

func (helper grpcRequestHelper) GetTenant(ctx context.Context) (string, error) {
	// Extract metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("could not extract metadata")
	}

	// Check if the X-Company-Id header is present
	companyID := md[auth.XTenantKey]
	if len(companyID) == 0 {
		return "", errors.New("could not extract company id")
	}

	log.Printf("Received X-Company-Id: %s", companyID[0])

	return companyID[0], nil
}

func NewGrpcRequestHelper(authenticator auth.ServiceAuthenticator) types.RequestHelper {
	return &grpcRequestHelper{
		authenticator: authenticator,
	}
}
