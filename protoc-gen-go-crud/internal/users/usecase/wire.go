package usecase

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserUseCase)