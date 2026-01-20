package handler

import (
	"context"

	pb "github.com/jne100/golang-service-layout/api"
	"github.com/jne100/golang-service-layout/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateItem(
	ctx context.Context,
	in *pb.CreateItemRequest,
) (*pb.CreateItemResponse, error) {
	if err := h.argsValidator.Validate(ctx,
		h.argsValidator.SaneSKU(in.Item.Sku),
		h.argsValidator.SaneItemName(in.Item.Name),
		h.argsValidator.PositiveInt32(in.Item.Quantity),
	); err != nil {
		return nil, err
	}

	err := h.ctrl.CreateItem(ctx, model.FromPbItem(in.Item))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create item: %v", err)
	}

	return &pb.CreateItemResponse{}, nil
}
