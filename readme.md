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

## 项目目录说明

> [info]
- core // 核心程序，modsec审计日志的结构体说明
- dao  // 数据库相关管理程序
- handles 
   - conf // 从yaml配置文件读取信息。
   - parser // 转化时间，转化告警消息的相关模块
   - rdfile  // 读取文件工具，测试时候使用过，当前读取的tail行，可舍弃 [弃用]
- logsource //本部分从文件流读取，已经在`main`中重写 [弃用]
   - file   // 从文本tail行读取，当然后面本目录可以增加从http/syslog等源的信息
- tests 测试相关组件的测试文件 [弃用]
- utils 工具模块
   - logging // 日志存储模块
   - path  // 读取文件路径的功能模块；当前未使用，因为`config.yaml`可以通过相对路径获取。
- config.yaml  //项目配置文件 
- main.go      //项目入口 
