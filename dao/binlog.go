package dao

import (
	"encoding/json"
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/spf13/viper"
	"time"
)

//startBinlog() 相关内容

type MyEventHandler struct {
	canal.DummyEventHandler
}

//OnRow 这个方法会被循环调用
func (h *MyEventHandler) OnRow(ev *canal.RowsEvent) error {

	rowData := make(map[string]interface{})
	for columnIndex, currColumn := range ev.Table.Columns {
		//字段名，字段的索引顺序，字段对应的值
		//row := fmt.Sprintf("%v %v %v\n", currColumn.Name, columnIndex, ev.Rows[len(ev.Rows)-1][columnIndex])
		rowData[currColumn.Name] = ev.Rows[len(ev.Rows)-1][columnIndex]
		//fmt.Println("currColumn.Name", currColumn.Name)
	}
	fmt.Println("Table:", ev.Table.Name)
	fmt.Println("rowData:", rowData)

	rowJson, err := json.Marshal(rowData)
	if err != nil {
		panic(err)
	}
	//把数据发给 Kafka
	err = KafkaMQ.Pub(ev.Table.Name, rowJson)
	if err != nil {
		panic(err)
	}
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func BinlogSyncInit() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%v:%v",viper.Get("mysql.host"), viper.Get("mysql.port"))
	cfg.User = fmt.Sprintf("%v", viper.Get("mysql.user"))
	cfg.Password = fmt.Sprintf("%v", viper.Get("mysql.password"))
	cfg.ServerID = 1
	cfg.ReadTimeout = 60 * time.Second

	cfg.Dump.TableDB = fmt.Sprintf("%v", viper.Get("mysql.database"))

	//这里放需要同步 binlog 的表名
	cfg.Dump.Tables = []string{"order", "category"}

	fmt.Println(cfg.Dump.Tables)

	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
