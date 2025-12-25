package main

import (
	"fmt"
	"server/internal/bootstrap"
	"server/internal/infra"
)

func main() {
	bootstrap.InitConfig()

	if err := bootstrap.MigrateDatabase(); err != nil {
		panic(err)
	}

	if err := bootstrap.InitDatabase(); err != nil {
		panic(err)
	}

	if err := bootstrap.InitRds(); err != nil {
		panic(err)
	}

	r := bootstrap.InitGin()
	addr := fmt.Sprintf("%s:%d", infra.AppConfig.App.Host, infra.AppConfig.App.Port)
	r.Run(addr)
}
