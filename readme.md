## golang 使用
- [go-json](https://www.cnblogs.com/yorkyang/p/8990570.html)
- [go代理](https://studygolang.com/articles/23826?fr=sidebar) 
   - `go env -w GOPROXY=https://goproxy.cn,direct`

## 2021-7-14 
- 开始测试golang的hi使用，包调用。


## 工作步骤
- 1. 从文件`tail`后, 将数据丢到一个信道中进行异步处理。
- 2. 信道处理来的数据，进行日期格式的转换，告警消息的转化，字段转化。
- 3. 每当转换的数据条目达到`<max>100`条时或者时间达到200，就把保存的数据都写入到数据库


## 2021-7-19 
```bash

docker exec -t mysql mysql -uroot -ptest@1q2w2e4R -e 'CREATE DATABASE `logger1` DEFAULT CHARACTER SE
T utf8 COLLATE utf8_general_ci;'

```
