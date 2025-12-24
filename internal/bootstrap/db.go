package bootstrap

import (
	"fmt"
	"server/internal/infra"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDatabase() error {
	dbCfg := AppConfig.Database
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%s", dbCfg.Dsn()),
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func InitDatabase() error {
	dbCfg := AppConfig.Database
	return infra.InitDB(dbCfg.Dsn())
}
