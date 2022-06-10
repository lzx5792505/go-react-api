package config

import (
	"liu/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

// 动态加载配置文件
type ConfigFunc func() map[string]interface{}

var ConfigFuncS map[string]ConfigFunc

func init() {
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()
	ConfigFuncS = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEnv(env)
	loadConfig()
}

func loadEnv(suffix string) {
	envPath := ".env"
	if len(suffix) > 0 {
		filepath := ".env." + suffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncS {
		viper.Set(name, fn())
	}
}

func Env(name string, val ...interface{}) interface{} {
	if len(val) > 0 {
		return internalGet(name, val[0])
	}
	return internalGet(name)
}

func Add(name string, configFn ConfigFunc) {
	ConfigFuncS[name] = configFn
}

func Get(path string, val ...interface{}) string {
	return GetString(path, val...)
}

func internalGet(path string, val ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(val) > 0 {
			return val[0]
		}
		return nil
	}
	return viper.Get(path)
}

func GetString(path string, val ...interface{}) string {
	return cast.ToString(internalGet(path, val...))
}

func GetInt(path string, val ...interface{}) int {
	return cast.ToInt(internalGet(path, val...))
}

func GetFloat64(path string, val ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, val...))
}

func GetInt64(path string, val ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, val...))
}

func GetUint(path string, val ...interface{}) uint {
	return cast.ToUint(internalGet(path, val...))
}

func GetBool(path string, val ...interface{}) bool {
	return cast.ToBool(internalGet(path, val...))
}

func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
