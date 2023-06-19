package deliveries

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewRoleService, NewAuthService)