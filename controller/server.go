package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"job3_canal/dao"
)

func StartBinlog() {
	fmt.Println("@")
	dao.InitKafka()
	fmt.Println("2")
	dao.BinlogSyncInit()
}

func StartEs() {
	dao.InitKafka()
	dao.InitES()
	dao.CategorySync()
}

func StartEsCommodity() {
	dao.InitKafka()
	dao.InitES()
	dao.CommoditySync()
}

func StartHttp() {
	//ES的一些业务需求，可以抽离出去优化
	dao.InitES()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/OrderSearch", OrderSearch)
	r.GET("/CategorySearch", CategorySearch)
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
