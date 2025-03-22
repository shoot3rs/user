package helpers

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/shoot3rs/user/internal/auth"
	"github.com/shoot3rs/user/internal/types"
	"google.golang.org/grpc/metadata"
	"log"
)

type contextHelper struct {
	authenticator auth.ServiceAuthenticator
}

func (helper contextHelper) GetUserClaimsFromRequest(request *connect.Request[connect.AnyRequest]) *auth.UserAuthClaims {
	request.Header().Get("Authorization")

	return nil
}

func (helper contextHelper) GetUserClaims(ctx context.Context) *auth.UserAuthClaims {
	userClaims := ctx.Value(auth.ContextKeyUser).(*auth.UserAuthClaims)
	return userClaims
}

func (helper contextHelper) GetTenant(ctx context.Context) (string, error) {
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

func NewContextHelper(authenticator auth.ServiceAuthenticator) types.ContextHelper {
	return &contextHelper{
		authenticator: authenticator,
	}
}
