package handler

import (
	"context"

	pb "github.com/jne100/golang-service-layout/api"
	"github.com/jne100/golang-service-layout/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetItem(
	ctx context.Context,
	in *pb.GetItemRequest,
) (*pb.GetItemResponse, error) {
	if err := h.argsValidator.Validate(ctx,
		h.argsValidator.SaneSKU(in.Sku),
	); err != nil {
		return nil, err
	}

	item, err := h.ctrl.GetItem(ctx, in.Sku)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get item: %v", err)
	}

	return &pb.GetItemResponse{
		Item: model.ToPbItem(item),
	}, nil
}
