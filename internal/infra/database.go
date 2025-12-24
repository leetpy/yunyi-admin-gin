package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	mysqlConfig := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}

	// opts := gorm.Config{
	// 	Logger: logger.New(NewWriter(general), logger.Config{
	// 		SlowThreshold: 200 * time.Millisecond,
	// 		LogLevel:      general.LogLevel(),
	// 		Colorful:      true,
	// 	}),
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   prefix,
	// 		SingularTable: singular,
	// 	},
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// }

	var err error
	if DB, err = gorm.Open(mysql.New(mysqlConfig)); err != nil {
		// global.G_LOG.Error("mysql connect failed", zap.Error(err))
	}

	return err
}
