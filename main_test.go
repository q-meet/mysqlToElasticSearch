package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"job3_canal/dao"
	"reflect"
	"testing"

	"github.com/olivere/elastic/v7"
)

func ConfigInit() {
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


func TestIndex(t *testing.T) {

	ConfigInit()

	dao.InitES()

	//短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("order_id", "hello")
	//条件查询
	//年龄大于30岁的
	boolQuery := elastic.NewBoolQuery()
	//字段相等
	//elastic.NewQueryStringQuery("last_name:Smith")
	// =查询 term实现单值精确匹配 terms就可以实现多值匹配
	boolQuery.Must(elastic.NewTermQuery("flag", "recommend"))
	// !=查询
	boolQuery.MustNot(elastic.NewMatchPhraseQuery("nickname", "company-crm"))
	// 过滤 范围查询 range就可以实现范围查询
	boolQuery.Filter(elastic.NewRangeQuery("pid").Gt(2))

	//短语搜索 搜索about字段中有 rock climbing
	boolQuery.Must(elastic.NewMatchPhraseQuery("diyname", "company"))

	// 或者查询
	// boolQuery.Should(elastic.NewTermQuery("pid", "6"), elastic.NewTermQuery("diyname", "test1"))

	res, err := dao.Client.Search("category").Query(boolQuery).Do(dao.Ctx)

	if err != nil {
		panic(res)
	}

	var typ dao.Category
	var ress []dao.Category
	// 通过Each方法，将es结果的json结构转换成struct对象
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		// 转换成Article对象
		t := item.(dao.Category)
		fmt.Printf("%#v\n\n", t)
		ress = append(ress, t)
	}

	fmt.Println()
	fmt.Println()
	jsonStr, _ := json.Marshal(ress)
	fmt.Println(string(jsonStr))

}

// func (r *SearchRequest) ToFilter() *EsSearch {
// 	var search EsSearch
// 	if len(r.Nickname) != 0 {
// 		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("nickname", r.Nickname))
// 	}
// 	if len(r.Phone) != 0 {
// 		search.ShouldQuery = append(search.ShouldQuery, elastic.NewTermsQuery("phone", r.Phone))
// 	}
// 	if len(r.Ancestral) != 0 {
// 		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("ancestral", r.Ancestral))
// 	}
// 	if len(r.Identity) != 0 {
// 		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("identity", r.Identity))
// 	}

// 	if search.Sorters == nil {
// 		search.Sorters = append(search.Sorters, elastic.NewFieldSort("create_time").Desc())
// 	}

// 	search.From = (r.Num - 1) * r.Size
// 	search.Size = r.Size
// 	return &search
// }

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}


//创建
func TestCreate(t *testing.T) {

	ConfigInit()

	dao.InitES()

	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err :=dao.Client.Index().
		Index("megacorp1").
		//Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
}