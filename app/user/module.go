package user

import "go.uber.org/fx"

var UserModule = fx.Module("user", fx.Provide(NewUserRepository))
