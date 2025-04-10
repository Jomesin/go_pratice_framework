package conf

import (
	logger "Go_workspace/utils/log"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"path"
	"sync"
)

type config struct {
	Base struct {
		Address string `ini:"address"`
	} `ini:"base"`

	Middleware struct {
		MaxConcurrentRequests int `ini:"max_concurrentRequests"`
	} `ini:"middleware"`
}

var (
	once sync.Once
	Cfg  config
)

func init() {
	// 获取当前运行环境
	envName, exists := os.LookupEnv("env")
	if !exists {
		// 若不存在,则赋值默认测试环境
		envName = "test"
	}
	// 获取根目录路径
	dir, err := os.Getwd()
	if err != nil {
		logger.Error(fmt.Sprintf("Fail to get wd path: %v", err))
		panic(err)
	}
	confPath := path.Join(dir, "conf")
	envPath := path.Join(confPath, envName+"_env.ini")
	err = ini.MapTo(&Cfg, envPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Fail to read file: %v", err))
		panic(err)
	}
}
