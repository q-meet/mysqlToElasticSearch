package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"job3_canal/controller"
	"log"
)

var s = flag.String("s", "http", "Input your server type")

func main() {
	configInit()

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	flag.Parse()
	switch *s {
	case "http":
		controller.StartHttp()
	case "binlog": //go run main.go -s=binlog
		controller.StartBinlog()
	case "esCommodity": //go run main.go -s=esCommodity
		controller.StartEsCommodity()
	case "es": //go run main.go -s=es
		controller.StartEs()
	default:
		controller.StartHttp()
	}
}

func configInit() {
	/*viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})*/

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
		} else {
			// 配置文件被找到，但产生了另外的错误
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

}
