//+ wireinject

package main

import (
	"github.com/tatsuo48/go_ddd/application/service"
	"github.com/tatsuo48/go_ddd/domain/model"
	"github.com/tatsuo48/go_ddd/infrastructure/repositry"

	"github.com/google/wire"
)

func initUserApplicationService(fileName string) *service.UserApplicationService {
	wire.Build(
		repositry.NewInMemoryUserRepositry,
		model.NewUserFactory,
		service.NewUserApplicationService,
	)
	return nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}
