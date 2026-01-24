package handler

import (
	"fmt"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/jne100/golang-service-layout/api"
	"github.com/jne100/golang-service-layout/internal/config"
	"github.com/jne100/golang-service-layout/internal/controller"
	"github.com/jne100/golang-service-layout/internal/handler/argsvalidator"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

const (
	MaxMsgSize = 16 * 1024 * 1024 // 16 MB
)

var Module = fx.Options(
	argsvalidator.Module,
	fx.Provide(NewServer),
	fx.Provide(NewHandler),
)

type ServerParams struct {
	fx.In
	InventoryServer pb.InventoryServiceServer
}

func NewServer(p ServerParams) *grpc.Server {
	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(MaxMsgSize),
		grpc.MaxSendMsgSize(MaxMsgSize),
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
		),
	)

	pb.RegisterInventoryServiceServer(s, p.InventoryServer)

	// Should call after done with registering
	grpc_prometheus.EnableHandlingTimeHistogram()

	return s
}

type HandlerParams struct {
	fx.In
	ArgsValidator argsvalidator.ArgsValidator
	Ctrl          controller.Controller
}

func NewHandler(p HandlerParams) pb.InventoryServiceServer {
	return &handler{
		argsValidator: p.ArgsValidator,
		ctrl:          p.Ctrl,
	}
}

type RunHandlerParams struct {
	fx.In
	Server *grpc.Server
	Cfg    config.InventoryConfig
}

func RunHandlerAsync(p RunHandlerParams) {
	go func() error {
		grpcEndpoint := fmt.Sprintf("0.0.0.0:%d", p.Cfg.Inbound.GrpcPort)
		l, err := net.Listen("tcp", grpcEndpoint)
		if err != nil {
			return err
		}
		return p.Server.Serve(l)
	}()
}
