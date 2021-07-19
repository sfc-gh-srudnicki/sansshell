package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/snowflakedb/unshelled/services"
)

// Serve wraps up buildServer in a succinct API for callers
func Serve(hostport string, policy string) error {
	lis, err := net.Listen("tcp", hostport)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s, err := buildServer(lis, policy)
	if err != nil {
		return err
	}

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

// buildServer creates a gRPC server, attaches the OPA policy interceptor,
// registers all of the imported Unshelled modules. Separating this from Serve
// primarily facilitates testing.
func buildServer(lis net.Listener, policy string) (*grpc.Server, error) {
	o, err := NewOPA(policy)
	if err != nil {
		return &grpc.Server{}, fmt.Errorf("NewOpa: %s", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(o.Authorize))
	for _, unshelledService := range services.ListServices() {
		unshelledService.Register(s)
	}
	return s, nil
}
