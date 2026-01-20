package model

import (
	"time"

	pb "github.com/jne100/golang-service-layout/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Item struct {
	Sku       string
	Name      string
	Quantity  int
	CreatedAt time.Time
}

func ToPbItem(in Item) *pb.Item {
	return &pb.Item{
		Sku:       in.Sku,
		Name:      in.Name,
		Quantity:  int32(in.Quantity),
		CreatedAt: timestamppb.New(in.CreatedAt),
	}
}

func FromPbItem(in *pb.Item) Item {
	return Item{
		Sku:       in.Sku,
		Name:      in.Name,
		Quantity:  int(in.Quantity),
		CreatedAt: in.CreatedAt.AsTime(),
	}
}
