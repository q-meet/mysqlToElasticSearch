package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"job3_canal/dao"
	"reflect"
)

/**
  可以根据先写 SQL ，然后用 DSL 转成 json
  POST _sql/translate
  {
    "query": "SELECT * FROM commodity where status=1 and name like '%测试%' order by id desc"
  }
*/

func CategorySearch(c *gin.Context) {
	boolQuery := elastic.NewBoolQuery()

	if name, ok := c.GetQuery("name"); ok {
		//短语搜索 搜索about字段中有 ..
		boolQuery.Must(elastic.NewMatchPhraseQuery("nickname", name))
	}

	//boolQuery.Must(elastic.NewTermQuery("flag", "recommend"))
	//matchPhraseQuery := elastic.NewWildcardQuery("name", name)
	res, err := dao.Client.Search("category").Query(boolQuery).Do(dao.Ctx)

	if err != nil {
		c.JSON(200, gin.H{
			"errno":  err.Error(),
			"search": "1",
		})
		return
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", res.TookInMillis, res.TotalHits())
	// 查询结果不为空，则遍历结果
	var typ dao.Category
	var resule []dao.Category
	// 通过Each方法，将es结果的json结构转换成struct对象
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		// 转换成Article对象
		t := item.(dao.Category)
		fmt.Printf("%#v\n", t)
		resule = append(resule, t)
	}

	if err != nil {
		c.JSON(200, gin.H{
			"errno": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": resule,
	})
	return
}

func OrderSearch(c *gin.Context) {
	if name, ok := c.GetQuery("name"); ok {

		//短语搜索 搜索about字段中有 rock climbing
		matchPhraseQuery := elastic.NewMatchPhraseQuery("order_id", name)
		//matchPhraseQuery := elastic.NewWildcardQuery("order_id", name)
		res, err := dao.Client.Search("order").Query(matchPhraseQuery).Do(dao.Ctx)

		if err != nil {
			c.JSON(200, gin.H{
				"errno": err.Error(),
				"msg":   "search err",
			})
			return
		}

		fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", res.TookInMillis, res.TotalHits())
		// 查询结果不为空，则遍历结果
		var typ dao.Order
		var resule []dao.Order
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
			// 转换成Article对象
			t := item.(dao.Order)
			fmt.Printf("%#v\n", t)
			resule = append(resule, t)
		}

		if err != nil {
			c.JSON(200, gin.H{
				"errno": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"data": resule,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "未传递参数",
	})
}
