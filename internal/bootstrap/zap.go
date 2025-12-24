package bootstrap

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// func InitZapLogger() error {
// 	logCfg := AppConfig.Zap

// 	if ok, _ := utils.PathExists(logCfg.Director); !ok { // 判断是否有Director文件夹
// 		fmt.Printf("create %v directory\n", logCfg.Director)
// 		_ = os.Mkdir(logCfg.Director, os.ModePerm)
// 	}

// 	levels := logCfg.Levels()
// 	length := len(levels)
// 	cores := make([]zapcore.Core, 0, length)
// 	for i := 0; i < length; i++ {
// 		core := internal.NewZapCore(levels[i])
// 		cores = append(cores, core)
// 	}
// 	Logger = zap.New(zapcore.NewTee(cores...))
// 	if logCfg.ShowLine {
// 		Logger = Logger.WithOptions(zap.AddCaller())
// 	}
// 	return nil
// }
