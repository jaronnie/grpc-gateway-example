package logx

import "github.com/spf13/viper"

// return true if key in config file
func viperContainsKey(key string) (ct bool) {
	keys := viper.GetViper().AllKeys()
	if len(keys) == 0 {
		return false
	}

	for _, k := range keys {
		if k == key {
			return true
		}
	}

	return false
}

// Set default string if viper config not set
func defaultViperString(key, def string) (value string) {
	if viperContainsKey(key) {
		return viper.GetString(key)
	}

	return def
}

// Set default int if viper config not set
func defaultViperInt(key string, def int) (value int) {
	if viperContainsKey(key) {
		return viper.GetInt(key)
	}

	return def
}

// Set default bool if viper config not set
func defaultViperBool(key string, def bool) (value bool) {
	if viperContainsKey(key) {
		return viper.GetBool(key)
	}

	return def
}
