package dao

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/jmoiron/sqlx"
	conf "golog/handles"
	"log"
	modsec "golog/core"
	rdfile "golog/handles"
	logging "golog/utils"
)

var Db *sqlx.DB

func init() {
	config := conf.GetConf("D:\\home\\projects\\golog\\config.yaml")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Mysql[0].Username,
		config.Mysql[0].Password,
		config.Mysql[0].Host,
		config.Mysql[0].Port,
		config.Mysql[0].Database,
		)

	database, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	Db = database
}


var schema = `
CREATE TABLE audit_modsec_v1 (
	event_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '时间戳',
	client_ip VARCHAR(50) DEFAULT '' COMMENT '客户端IP',
	client_port INTEGER DEFAULT 1180 COMMENT '客户端端口',
	host_ip VARCHAR(50) DEFAULT '' COMMENT '服务端IP',
	host_port INTEGER DEFAULT 1180 COMMENT '服务端端口',
	unique_id VARCHAR(100) NOT NULL,
	server_id VARCHAR(100) DEFAULT '' COMMENT '服务端ID',
	method VARCHAR(10) DEFAULT 'GET' COMMENT '请求方法',
	uri VARCHAR(255) DEFAULT '' COMMENT 'URI',
	args text DEFAULT NULL COMMENT '请求参数',
	user_agent text DEFAULT NULL COMMENT '请求客户端',
	http_version VARCHAR(20) DEFAULT '' COMMENT 'HTTP版本',
    req_headers text DEFAULT NULL COMMENT '请求头',
    res_headers text DEFAULT NULL COMMENT '响应头',
    req_body text DEFAULT NULL COMMENT '请求体',
    res_body text DEFAULT NULL COMMENT '响应体',
    res_c_type VARCHAR(150) DEFAULT '' COMMENT '响应文件类型',
    status INTEGER DEFAULT 100 COMMENT '状态码',
    msg VARCHAR(255) DEFAULT '自定义告警' COMMENT '告警消息',
    category VARCHAR(255) DEFAULT '策略不合规' COMMENT '告警消息',
    waf_info text DEFAULT NULL COMMENT 'WAF信息',
	PRIMARY KEY (unique_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var drop_schema = `drop table if exists audit_modsec_v1;`


func DropTable(){
	Db.MustExec(drop_schema)
	log.Println("Drop OK")
}

func ReCreateTable(){
	DropTable()

	Db.MustExec(schema)
	log.Println("Create OK")
}

func InsertMutiData(){

	var alis []modsec.AuditLogItem
	lines := rdfile.ReadFile()
	for _, line  := range lines {
		res := modsec.ParseSingleLine(line)
		if res.UniqId != "" {
			alis = append(alis, res)
		}else {
			fmt.Println("Parse Error...")
		}
		//fmt.Println(res)
	}
	BulkInsert(alis)
}

func BulkInsert(alis []modsec.AuditLogItem ){
	var query = `INSERT INTO audit_modsec_v1 
		(event_time, client_ip, client_port, host_ip, 
		host_port, unique_id, server_id, method, 
		uri, args, user_agent, http_version, 
		req_headers, res_headers, req_body, res_body, 
		res_c_type, status, msg, category, waf_info)
        VALUES (:event_time, :client_ip, :client_port, :host_ip, 
        :host_port, :unique_id, :server_id, :method, 
        :uri, :args, :user_agent, :http_version, 
        :req_headers, :res_headers, :req_body, :res_body, 
        :res_c_type, :status, :msg, :category, :waf_info)`
	_, err := Db.NamedExec(query, alis)

	if err != nil {
		logging.Error.Fatalln(err)
		return
	}

	logging.Warning.Println("Insert OK...")
}