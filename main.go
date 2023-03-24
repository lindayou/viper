package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var stringFlag = pflag.String("stringflag", "stringflag", "string flag usage")
var stringpFlag = pflag.StringP("stringpflag", "s", "stringpflag", "stringp flag usage")
var intFlag int
var boolFlag bool

func init() {
	pflag.IntVarP(&intFlag, "intflag", "i", 0, "int flag usage")
	pflag.BoolVar(&boolFlag, "boolflag", false, "bool flag usage")
}

func main() {

	v := viper.New()
	v.AddConfigPath(".")         // 添加配置文件搜索路径，点号为当前目录
	v.AddConfigPath("./configs") // 添加多个搜索目录
	v.SetConfigType("yaml")      // 如果配置文件没有后缀，可以不用配置
	v.SetConfigName("app.yml")   // 文件名，没有后缀

	// v.SetConfigFile("configs/app.yml")

	// 读取配置文件
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())

	}

	// 通过.号来区分不同层级，来获取配置值
	log.Printf("app.mode=%s\n", v.Get("app.mode"))
	log.Printf("db.mysql.url=%s\n", v.Get("db.mysql.url"))
	log.Printf("db.redis.host=%s\n", v.GetString("db.redis.host"))
	log.Printf("db.redis.port=%d\n", v.GetInt("db.redis.port"))

	// 使用Sub获取子配置，然后获取配置值
	v2 := v.Sub("db")
	log.Printf("db.mysql.url:%s\n", v2.Sub("mysql").GetString("url"))
	log.Printf("db.redis.host:%s\n", v2.Sub("redis").GetString("host"))
	log.Printf("db.redis.port:%s\n", v2.Sub("redis").GetString("port"))
	return
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")         // 添加配置文件搜索路径，点号为当前目录
	v.AddConfigPath("./configs") // 添加多个搜索目录
	v.SetConfigType("yaml")      // 如果配置文件没有后缀，可以不用配置
	v.SetConfigName("app.yml")   // 文件名，没有后缀

	// v.SetConfigFile("configs/app.yml")

	// 读取配置文件
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}
	return v, nil
}
