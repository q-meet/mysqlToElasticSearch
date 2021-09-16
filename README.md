# 从Mysql到Elasticsearch

## 流程

MySQL -> Binlog  -> kafka -> consumer -> elasticsearchEs —> query -> keyword -> 搜索结果

1. 导入sql [database.sql](./database.sql)

2. 配置文件 配置 数据库 kafka es 相关配置

以下示列配置

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: root
  database: test
es:
  host: http://127.0.0.1
  port: 9200
  user:
  password:
kafka:
  host: localhost
  port: 9092
  topic: kafka_log
```

3. binlog To kafka

```shell
go run main.go -s=binlog
```

4. kafka To elasticsearchEs

```shell
go run main.go -s=es
```

5. http 展示

```shell
go run main.go -s=http
```

访问url:
[http://127.0.0.1:8088/CategorySearch](http://127.0.0.1:8088/CategorySearch)

筛选 做了nicename的精准匹配
[http://127.0.0.1:8088/CategorySearch?name=news](http://127.0.0.1:8088/CategorySearch?name=news)
