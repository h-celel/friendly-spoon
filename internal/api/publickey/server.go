package publickey

import (
	"context"
	"fmt"
	"github.com/h-celel/friendly-spoon/internal/config"
	"github.com/h-celel/friendly-spoon/internal/service/publickey"
	"github.com/h-celel/friendly-spoon/proto/spoon/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type handler struct {
	auth.UnimplementedAuthenticationAuditServiceServer
	publicKeyService publickey.Service
}

func (s *handler) GetPublicKey(_ context.Context, _ *auth.GetPublicKeyRequest) (*auth.GetPublicKeyResponse, error) {
	return &auth.GetPublicKeyResponse{
		Response: &auth.GetPublicKeyResponse_ResponseBody{
			ResponseBody: &auth.GetPublicKeyResponse_GetPublicKeyResponseBody{
				PublicKey: s.publicKeyService.Get(),
			},
		},
	}, nil
}

func newHandler(publicKeyService publickey.Service) *handler {
	return &handler{
		publicKeyService: publicKeyService,
	}
}

func Init(ctx context.Context, cancel context.CancelFunc, env config.Environment) {
	publicKeyService := publickey.NewService(ctx, fmt.Sprintf("https://%s/pem", env.Auth0Domain))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.GRPCInternalPort))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	handler := newHandler(publicKeyService)

	go func() {
		defer cancel()
		auth.RegisterAuthenticationAuditServiceServer(server, handler)
		reflection.Register(server)
		err = server.Serve(lis)
		if err != nil {
			log.Println(err)
		}
	}()
}
