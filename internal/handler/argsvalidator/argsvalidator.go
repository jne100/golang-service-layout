package argsvalidator

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArgValidator func(ctx context.Context) error

type ArgsValidator interface {
	Validate(ctx context.Context, validators ...ArgValidator) error
	SaneSKU(id string) ArgValidator
	SaneItemName(name string) ArgValidator
	PositiveInt32(x int32) ArgValidator
}

type argsValidator struct {
}

func (v *argsValidator) Validate(ctx context.Context, validators ...ArgValidator) error {
	for _, validator := range validators {
		if err := validator(ctx); err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
		}
	}
	return nil
}

func (v *argsValidator) SaneSKU(sku string) ArgValidator {
	return func(ctx context.Context) error {
		if x, err := strconv.Atoi(sku); err != nil {
			return fmt.Errorf("invalid sku: %w", err)
		} else if x < 0 {
			return fmt.Errorf("invalid sku: sku %q is negative", sku)
		}
		return nil
	}
}

func (v *argsValidator) SaneItemName(name string) ArgValidator {
	return func(ctx context.Context) error {
		if len(name) > 128 {
			return fmt.Errorf("invalid item name: length %d exceeds maximum of 128", len(name))
		}
		return nil
	}
}

func (v *argsValidator) PositiveInt32(x int32) ArgValidator {
	return func(ctx context.Context) error {
		if x <= 0 {
			return fmt.Errorf("invalid integer: must be positive")
		}
		return nil
	}
}
