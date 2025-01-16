package middlewares

import (
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"context"
	"errors"
	"fmt"
	"github.com/rs/cors"
	"github.com/shoot3rs/user/internal/auth"
	"github.com/shoot3rs/user/internal/routes"
	"github.com/shoot3rs/user/internal/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"slices"
	"time"
)

type grpcAuthMiddleware struct{}

func (middleware *grpcAuthMiddleware) CorsMiddleware(h http.Handler) http.Handler {
	var allowedHeaders, allowedOrigins, allowedMethods []string
	allowedHeaders = append(allowedHeaders, connectcors.AllowedHeaders()...)
	allowedOrigins = append(allowedOrigins, "*")
	allowedMethods = append(allowedMethods, connectcors.AllowedMethods()...)
	allowedMethods = append(allowedMethods, "OPTIONS")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedHeaders: allowedHeaders,
		AllowedMethods: allowedMethods,
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	return corsMiddleware.Handler(h)
}

func (middleware *grpcAuthMiddleware) UnaryTokenInterceptor(authenticator auth.ServiceAuthenticator) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Extract the full method name from the request
			fullMethod := req.Spec().Procedure
			log.Println("Request FullMethod ::::: |", fullMethod)

			// If the method is public, skip authentication
			if slices.Contains(routes.PublicRoutes, fullMethod) {
				log.Println("FullMethod exists in public routes ::::: |", fullMethod)
				return next(ctx, req)
			}

			// Otherwise, apply the authentication middleware
			token, err := authenticator.ExtractHeaderToken(req)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(fmt.Sprintf("missing or invalid token: %v", err)))
			}

			// Validate the token
			idToken, err := authenticator.GetVerifier().Verify(ctx, token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(fmt.Sprintf("invalid token: %v", err)))
			}

			// Add user info to context
			claims := new(auth.UserAuthClaims)
			if err := idToken.Claims(claims); err != nil {
				return nil, status.Error(codes.Internal, fmt.Sprintf("failed to parse token claims: %v", err))
			}

			// Add the claims to the context
			newCtx := context.WithValue(ctx, auth.ContextKeyUser, claims)

			// Proceed with the handler
			return next(newCtx, req)
		}
	}
	return interceptor
}

func (middleware *grpcAuthMiddleware) UnaryTenantHeaderInterceptor(authenticator auth.ServiceAuthenticator) connect.UnaryFunc {
	//TODO implement me
	panic("implement me")
}

// wrappedServerStream wraps grpc.ServerStream to allow us to modify the context
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

// Context returns the modified context with the claims
func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

func sanitizeRequest(req interface{}) interface{} {
	// Implement logic to remove or mask sensitive fields
	return req
}

func (middleware *grpcAuthMiddleware) LoggingUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		// Sanitize the request to remove any sensitive data
		sanitizedReq := sanitizeRequest(req)

		log.Printf("gRPC Method: %s, Request: %+v", info.FullMethod, sanitizedReq)

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		if err != nil {
			log.Printf("gRPC Method: %s, Error: %v, Duration: %s", info.FullMethod, err, duration)
		} else {
			log.Printf("gRPC Method: %s, Response: %+v, Duration: %s", info.FullMethod, resp, duration)
		}

		return resp, err
	}
}

// TokenValidationUnaryInterceptor selectively applies authentication only to private methods
func (middleware *grpcAuthMiddleware) TokenValidationUnaryInterceptor(authenticator auth.ServiceAuthenticator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// If the method is public, skip authentication
		log.Println("Checking for public routes :::: |", routes.PublicRoutes)
		if slices.Contains(routes.PublicRoutes, info.FullMethod) {
			return handler(ctx, req)
		}

		// Otherwise, apply the authentication middleware
		token, err := authenticator.ExtractToken(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(fmt.Sprintf("missing or invalid token: %v", err)))
		}

		// Validate the token
		idToken, err := authenticator.GetVerifier().Verify(ctx, token)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New(fmt.Sprintf("invalid token: %v", err)))
		}

		// Add user info to context
		claims := new(auth.UserAuthClaims)
		if err := idToken.Claims(claims); err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to parse token claims: %v", err))
		}

		// Add the claims to the context
		newCtx := context.WithValue(ctx, auth.ContextKeyUser, claims)

		// Proceed with the handler
		return handler(newCtx, req)
	}
}

// StreamInterceptor selectively applies authentication only to private methods
func (middleware *grpcAuthMiddleware) StreamInterceptor(authenticator auth.ServiceAuthenticator) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		methodName := info.FullMethod

		// Check if the method is public
		if slices.Contains(routes.PublicRoutes, methodName) {
			log.Printf("Public method: %s", methodName)
			return handler(srv, ss) // No auth required, proceed
		}

		// If the method is private, check for authentication
		log.Printf("Private method: %s", methodName)

		// Extract metadata from the stream
		ctx := ss.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return connect.NewError(connect.CodeUnauthenticated, errors.New("missing metadata"))
		}

		// Check for the authorization token in the metadata
		tokens := md["authorization"]
		if len(tokens) == 0 {
			return connect.NewError(connect.CodeUnauthenticated, errors.New("invalid or missing token"))
		}

		// Remove the "Bearer " prefix from the token if it's present
		rawToken := tokens[0]
		if len(rawToken) > 7 && rawToken[:7] == "Bearer " {
			rawToken = rawToken[7:]
		}

		// Verify the token using the OIDC verifier
		idToken, err := authenticator.GetVerifier().Verify(ctx, rawToken)
		if err != nil {
			return connect.NewError(connect.CodeUnauthenticated, errors.New(fmt.Sprintf("invalid token: %v", err)))
		}

		// Optionally, check additional claims, such as the audience or expiration
		log.Printf("Token is valid, subject: %s, expiry: %v", idToken.Subject, idToken.Expiry)

		// Parse the token claims into a map
		// Add user info to context
		claims := new(auth.UserAuthClaims)
		if err := idToken.Claims(claims); err != nil {
			return status.Error(codes.Internal, "failed to parse token claims")
		}

		//log.Println("UserClaimsKey: ", claims)
		// Add the claims to the context
		newCtx := context.WithValue(ss.Context(), auth.ContextKeyUser, claims)

		// Wrap the stream with the new context containing claims
		wrappedStream := &wrappedServerStream{
			ServerStream: ss,
			ctx:          newCtx,
		}

		// Proceed with the stream handler
		return handler(srv, wrappedStream)
	}
}

// TenantHeaderInterceptor is a gRPC interceptor that checks for the presence of the X-Company-Id header
func (middleware *grpcAuthMiddleware) TenantHeaderInterceptor(excludedMethods []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		// Check if the current method is in the excluded methods list
		for _, method := range excludedMethods {
			if info.FullMethod == method {
				log.Printf("Skipping HeaderInterceptor for method: %s", info.FullMethod)
				return handler(ctx, req)
			}
		}

		// Extract metadata from context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
		}

		// Check if the X-Tenant-Id header is present
		tenant := md[auth.XTenantKey]
		if len(tenant) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "missing required X-Tenant-Id header")
		}

		log.Printf("Received X-Tenant-Id: %s", tenant[0])

		// Continue to the next handler
		return handler(ctx, req)
	}
}

func New() types.GRPCAuthMiddleware {
	return &grpcAuthMiddleware{}
}
