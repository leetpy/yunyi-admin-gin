package main

import (
	"fmt"
	"server/internal/bootstrap"
)

func main() {
	bootstrap.InitConfig()
	r := bootstrap.InitGin()
	addr := fmt.Sprintf("%s:%d", bootstrap.AppConfig.App.Host, bootstrap.AppConfig.App.Port)
	r.Run(addr)
}
