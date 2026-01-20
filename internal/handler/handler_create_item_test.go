package handler

import (
	"context"
	"errors"
	"testing"

	pb "github.com/jne100/golang-service-layout/api"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_CreateItem(t *testing.T) {
	t.Run("returns success", func(t *testing.T) {
		// Given: mocked handler
		h, mocks := newTestHandler(t)

		// Expect: controller returns success
		mocks.ctrlMock.EXPECT().CreateItem(context.Background(), gomock.Any()).Return(nil)

		// When: call CreateItem
		_, err := h.CreateItem(context.Background(), &pb.CreateItemRequest{
			Item: &pb.Item{Sku: "123"},
		})

		// Then: no error
		require.NoError(t, err)
	})

	t.Run("returns error", func(t *testing.T) {
		// Given: mocked handler
		h, mocks := newTestHandler(t)

		// Expect: controller returns error
		mocks.ctrlMock.EXPECT().
			CreateItem(context.Background(), gomock.Any()).
			Return(errors.New("failed to create"))

		// When: call CreateItem
		_, err := h.CreateItem(context.Background(), &pb.CreateItemRequest{
			Item: &pb.Item{Sku: "123"},
		})

		// Then: error is propagated
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})
}
