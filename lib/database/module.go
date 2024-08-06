package database

import (
	"go.uber.org/fx"
)

var DBModule = fx.Module(
	"db",
	fx.Provide(
		NewDB, // 함수를 바로 제공
	), fx.Invoke(func(db *Database) {}),
)
