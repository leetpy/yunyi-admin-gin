package main

import (
	"fmt"
	"server/internal/bootstrap"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	bootstrap.InitConfig()

	dbCfg := bootstrap.AppConfig.Database
	m, err := migrate.New(
		"file://migrations",
		dbCfg.Dsn(),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	r := bootstrap.InitGin()
	addr := fmt.Sprintf("%s:%d", bootstrap.AppConfig.App.Host, bootstrap.AppConfig.App.Port)
	r.Run(addr)
}
