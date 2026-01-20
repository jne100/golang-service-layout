package controller

import (
	"go.uber.org/fx"

	"github.com/jne100/golang-service-layout/internal/repository"
)

var Module = fx.Options(
	fx.Provide(NewController),
)

type Params struct {
	fx.In
	Repo repository.Repository
}

func NewController(p Params) Controller {
	return &controller{
		repo: p.Repo,
	}
}
