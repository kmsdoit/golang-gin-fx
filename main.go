package main

import (
	"go-server/app/server"
	"go-server/app/user"
	"go-server/lib/database"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		database.DBModule,
		user.UserModule,
		fx.Provide(server.NewServer),
		fx.Invoke(
			func(*http.Server) {},
		)).Run()
}
