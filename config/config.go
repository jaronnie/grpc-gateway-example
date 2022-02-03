package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/jaronnie/grpc-gateway-example/util"
	"github.com/spf13/viper"
)

const (
	defaultFilePath = "./conf/cfg.toml"
)

// CfgManager define config manager
type CfgManager struct{}

// NewConfigManager new config manager
func NewConfigManager(name string) *CfgManager {
	cm := &CfgManager{}
	return cm
}

// NewConfig new Configure
func (cm *CfgManager) NewConfig() error {
	return cm.newLocalConfig()
}

// newLocalConfig return local configure instance
func (cm *CfgManager) newLocalConfig() (err error) {

	// set local file path
	viper.SetConfigFile(defaultFilePath)
	ext, err := util.Ext(defaultFilePath, viper.SupportedExts...)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType(ext)

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// watch config and set onchange function
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO: local file change
		fmt.Println(fmt.Sprintf("Config file changed: %s\n", e.Name))
	})
	return
}
