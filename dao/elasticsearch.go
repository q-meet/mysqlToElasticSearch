package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

//dao/elasticsearch.go

var Client *elastic.Client
var Ctx context.Context
var Err error

// 索引mapping定义，这里仿微博消息结构定义
const commodity = `
mappings" : {
  "properties" : {
	"amount" : {
	  "type" : "long"
	},
	"create_time" : {
	  "type" : "text",
	  "fields" : {
		"keyword" : {
		  "type" : "keyword",
		  "ignore_above" : 256
		}
	  }
	},
	"id" : {
	  "type" : "long"
	},
	"order_id" : {
	  "type" : "text",
	  "fields" : {
		"keyword" : {
		  "type" : "keyword",
		  "ignore_above" : 256
		}
	  }
	}
  }
}`

func InitES() {
	// 创建client
	Client, Err = elastic.NewClient(
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(fmt.Sprintf("%v:%v", viper.Get("es.host"), viper.Get("es.port"))),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth(fmt.Sprintf("%v", viper.Get("es.user")), fmt.Sprintf("%s", viper.Get("es.password"))),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if Err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", Err)
	} else {
		fmt.Println("连接成功")
	}

	// 执行ES请求需要提供一个上下文对象
	Ctx = context.Background()

}

func CategorySync() {
	//订阅 nats 中的数据
	err := KafkaMQ.Sub("category", func(data []byte) error {
		//拿到数据 放到 es 中
		var category Category
		err := json.Unmarshal(data, &category)
		if err != nil {
			panic(err)
		}

		fmt.Println(category)

		// 使用client创建一个新的文档
		_, err = Client.Index().
			Index("category"). // 设置索引名称
			Id(strconv.Itoa(int(category.Id))). // 设置文档id
			BodyJson(category). // 指定前面声明的内容
			Do(Ctx) // 执行请求，需要传入一个上下文对象
		return err
	})
	if err != nil {
		panic(err)
	}
}

func CommoditySync() {
	// 首先检测下索引是否存在
	exists, err := Client.IndexExists("commodity").Do(Ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// 索引不存在，则创建一个
		_, err := Client.CreateIndex("commodity").BodyString(commodity).Do(Ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	//订阅 nats 中的数据
	err = KafkaMQ.Sub("order", func(data []byte) error {
		//拿到数据 放到 es 中
		var order Order
		err := json.Unmarshal(data, &order)
		if err != nil {
			panic(err)
		}
		fmt.Println(order)

		// 使用client创建一个新的文档
		_, err = Client.Index().
			Index("commodity"). // 设置索引名称
			Id(strconv.Itoa(int(order.Id))). // 设置文档id
			BodyJson(order). // 指定前面声明的微博内容
			Do(Ctx) // 执行请求，需要传入一个上下文对象

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
