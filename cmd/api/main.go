package main

import (
	"fmt"
	"server/internal/bootstrap"
)

func main() {
	bootstrap.InitConfig()

	if err := bootstrap.MigrateDatabase(); err != nil {
		panic(err)
	}

	if err := bootstrap.InitDatabase(); err != nil {
		panic(err)
	}

	r := bootstrap.InitGin()
	addr := fmt.Sprintf("%s:%d", bootstrap.AppConfig.App.Host, bootstrap.AppConfig.App.Port)
	r.Run(addr)
}
