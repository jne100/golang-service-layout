package handler

import (
	"context"
	"errors"
	"testing"

	pb "github.com/jne100/golang-service-layout/api"
	"github.com/jne100/golang-service-layout/internal/model"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GetItem(t *testing.T) {
	t.Run("returns success", func(t *testing.T) {
		// Given: mocked handler
		h, mocks := newTestHandler(t)

		// Expect: controller returns success
		item := model.Item{Sku: "123"}
		mocks.ctrlMock.EXPECT().GetItem(context.Background(), "123").Return(item, nil)

		// When: call GetItem
		resp, err := h.GetItem(context.Background(), &pb.GetItemRequest{Sku: "123"})

		// Then: no error
		require.NoError(t, err)
		require.Equal(t, "123", resp.Item.Sku)
	})

	t.Run("returns error", func(t *testing.T) {
		// Given: mocked handler
		h, mocks := newTestHandler(t)

		// Expect: controller returns error
		mocks.ctrlMock.EXPECT().
			GetItem(context.Background(), "123").
			Return(model.Item{}, errors.New("failed to get"))

		// When: call GetItem
		_, err := h.GetItem(context.Background(), &pb.GetItemRequest{Sku: "123"})

		// Then: error is propagated
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})
}
