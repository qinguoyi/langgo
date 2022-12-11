package bootstrap

import (
	"StorageProxy/config"
	"StorageProxy/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"sync"
)

var (
	configPath string
	rootPath   = utils.RootPath()
	lgConfig   = new(LangGoConfig)
)

// LangGoConfig 自定义Log
type LangGoConfig struct {
	Conf *config.Configuration
	Once *sync.Once
}

// newLangGoConfig .
func newLangGoConfig() *LangGoConfig {
	return &LangGoConfig{
		Conf: &config.Configuration{},
		Once: &sync.Once{},
	}
}

// NewConfig 初始化配置对象
func NewConfig() *config.Configuration {
	if lgConfig.Conf != nil {
		return lgConfig.Conf
	} else {
		lgConfig = newLangGoConfig()
		lgConfig.initLangGoConfig()
		return lgConfig.Conf
	}
}

// InitLangGoConfig 初始化日志
func (lg *LangGoConfig) initLangGoConfig() {
	lg.Once.Do(
		func() {
			initConfig(lg.Conf)
		},
	)
}

func initConfig(conf *config.Configuration) {
	pflag.StringVarP(&configPath, "conf", "", filepath.Join(rootPath, "go-storage-proxy/conf", "config.yaml"),
		"config path, eg: --conf config.yaml")
	if !filepath.IsAbs(configPath) {
		configPath = filepath.Join(rootPath, "conf", configPath)
	}

	//lgLogger.Logger.Info("load config:" + configPath)
	fmt.Println("load config:" + configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		//lgLogger.Logger.Error("read config failed: ", zap.String("err", err.Error()))
		fmt.Println("read config failed: ", zap.String("err", err.Error()))
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		//lgLogger.Logger.Error("config parse failed: ", zap.String("err", err.Error()))
		fmt.Println("config parse failed: ", zap.String("err", err.Error()))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//lgLogger.Logger.Info("", zap.String("config file changed:", in.Name))
		fmt.Println("", zap.String("config file changed:", in.Name))
		defer func() {
			if err := recover(); err != nil {
				//lgLogger.Logger.Error("config file changed err:", zap.Any("err", err))
				fmt.Println("config file changed err:", zap.Any("err", err))
			}
		}()
		if err := v.Unmarshal(&conf); err != nil {
			//lgLogger.Logger.Error("config parse failed: ", zap.String("err", err.Error()))
			fmt.Println("config parse failed: ", zap.String("err", err.Error()))
		}
	})
	lgConfig.Conf = conf
}
