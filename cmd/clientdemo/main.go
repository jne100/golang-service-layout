package main

import (
	"context"
	"time"

	inventorypb "github.com/jne100/golang-service-layout/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("grpc dial failed", zap.Error(err))
	}
	defer conn.Close()

	client := inventorypb.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.CreateItem(ctx, &inventorypb.CreateItemRequest{
		Item: &inventorypb.Item{
			Sku:       "123",
			Name:      "demo-item",
			Quantity:  1,
			CreatedAt: timestamppb.Now(),
		},
	})
	if err != nil {
		logger.Fatal("create item failed", zap.Error(err))
	}

	resp, err := client.GetItem(ctx, &inventorypb.GetItemRequest{Sku: "123"})
	if err != nil {
		logger.Fatal("get item failed", zap.Error(err))
	}

	logger.Info("item received",
		zap.String("sku", resp.Item.Sku),
		zap.String("name", resp.Item.Name),
		zap.Int32("quantity", resp.Item.Quantity),
	)
}
