package handler

import (
	"context"
	"testing"

	ctrlmocks "github.com/jne100/golang-service-layout/internal/controller/mocks"
	argsvalidatormocks "github.com/jne100/golang-service-layout/internal/handler/argsvalidator/mocks"
	"go.uber.org/mock/gomock"
)

type handlerMocks struct {
	ctrl          *gomock.Controller
	ctrlMock      *ctrlmocks.MockController
	argsValidator *argsvalidatormocks.MockArgsValidator
}

func newTestHandler(t *testing.T) (handler, handlerMocks) {
	ctrl := gomock.NewController(t)

	mocks := handlerMocks{
		ctrl:          ctrl,
		ctrlMock:      ctrlmocks.NewMockController(ctrl),
		argsValidator: argsvalidatormocks.NewMockArgsValidator(ctrl),
	}

	alwaysValidFn := func(context.Context) error { return nil }
	mocks.argsValidator.EXPECT().SaneSKU(gomock.Any()).AnyTimes().Return(alwaysValidFn)
	mocks.argsValidator.EXPECT().SaneItemName(gomock.Any()).AnyTimes().Return(alwaysValidFn)
	mocks.argsValidator.EXPECT().PositiveInt32(gomock.Any()).AnyTimes().Return(alwaysValidFn)
	mocks.argsValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return handler{
		ctrl:          mocks.ctrlMock,
		argsValidator: mocks.argsValidator,
	}, mocks
}
