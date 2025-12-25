package bootstrap

import (
	"flag"
	"fmt"
	"server/internal/config"
	"server/internal/infra"

	"github.com/spf13/viper"
)

func InitConfig() {
	configFile := getConfigPath()
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	infra.AppConfig = &cfg
}

func getConfigPath() (config string) {
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config != "" { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}

	return "etc/config.yaml"
}
