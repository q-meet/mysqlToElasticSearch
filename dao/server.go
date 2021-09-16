package dao

type MsgHandler func(n []byte) error

type Order struct {
	Id         int32   `json:"id"`
	OrderId    string  `json:"order_id"`
	Amount     float32 `json:"amount"`
	CreateTime string  `json:"create_time"`
}

type Category struct {
	Id          int32       `json:"id"`
	Pid         int32       `json:"pid"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	Nickname    string      `json:"nickname"`
	Flag        interface{} `json:"flag"`
	Image       string      `json:"image"`
	Keywords    string      `json:"keywords"`
	Description string      `json:"description"`
	Diyname     string      `json:"diyname"`
	Createtime  int32       `json:"createtime"`
	Updatetime  int32       `json:"updatetime"`
	Weigh       int32       `json:"weigh"`
	Status      string      `json:"status"`
}

type MessageQueue interface {
	Pub(key string, data []byte) error
	Sub(key string, mh MsgHandler) error
}

type KafkaMessageQueue interface {
	Pub(key string, data []byte) error
	Sub(key string, mh MsgHandler) error
	Clear() error
}
