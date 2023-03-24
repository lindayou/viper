package test

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	v, err := InitConfig()
	if err != nil {
		log.Fatalf("viper读取失败, error:%+v\n", err)
	}

	// 监听到文件变化后的回调
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.Get("db.redis.passwd"))
	})

	v.WatchConfig()

	// 阻塞进程退出
	time.Sleep(time.Duration(1000000) * time.Second)
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")          // 添加配置文件搜索路径，点号为当前目录
	v.AddConfigPath("../configs") // 添加多个搜索目录
	v.SetConfigType("yaml")       // 如果配置文件没有后缀，可以不用配置
	v.SetConfigName("app.yml")    // 文件名，没有后缀

	// v.SetConfigFile("configs/app.yml")

	// 读取配置文件
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}
	return v, nil
}
