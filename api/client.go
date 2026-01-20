package api

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

const (
	MaxMsgSize = 16 * 1024 * 1024 // 16 MB
)

var (
	DefaultConnectParams = grpc.ConnectParams{
		Backoff:           DefaultBackoff,
		MinConnectTimeout: DefaultMinConnectTimeout,
	}

	DefaultBackoff = backoff.Config{
		BaseDelay:  1 * time.Second,
		Multiplier: 1.6,
		Jitter:     0.2,
		MaxDelay:   10 * time.Second,
	}

	DefaultMinConnectTimeout = 5 * time.Second
)

// NewInventoryClient creates new inventory client
// Should close returned connection using conn.Close()
func NewInventoryClient(endpoint string) (InventoryServiceClient, error) {
	conn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure(),
		grpc.WithConnectParams(DefaultConnectParams),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(MaxMsgSize),
			grpc.MaxCallSendMsgSize(MaxMsgSize),
		),
	)
	if err != nil {
		return nil, err
	}
	return NewInventoryServiceClient(conn), nil
}
