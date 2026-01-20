package argsvalidator

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewArgsValidator),
)

type Params struct {
	fx.In
}

func NewArgsValidator(p Params) ArgsValidator {
	return &argsValidator{}
}
