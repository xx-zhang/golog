package dao

import (
	"fmt"
	"github.com/tidwall/gjson"
	rdfile "golog/handles"
	"strconv"

	"time"
)

// 定义 Modscurity 输出日志的结构体

type AuditLogItem struct {
	Timestamp time.Time
	ClientIp string
	ClientPort int
	UniqId string
	ServerID string
	HostIp string
	HostPort int
	SessionNum int
	RequestMethod string
	RequestUri string
	RequestBody string
	HttpVersion string
	UserAgent string
	RequestPayload string
	RequestHeaders string
	ResponseBody string
	ResponseStatus int
	ResponseHeaders string
	ResponseContentType string 
	AuditMsg string 
	WAFInfo string
}


const TIME_LAYOUT = "Thu Jul  1 04:03:29 2021"
const C_TIME_LAYLOUT = "2006-01-02 15:04:05"

func GetRandLine(){
	lines := rdfile.ReadFile()
	//fmt.Println(lines[0:1])
	//fmt.Println("-------------------")
	//rand.Seed(time.Now().UnixNano())
	//randIndex := rand.Intn(30)
	//randLine := lines[randIndex:randIndex+1]
	for _, line  := range lines {
		auditItem := &AuditLogItem{}
		gjson.ForEachLine(line, func(line gjson.Result) bool{
			transaction := line.Get("transaction")
			auditItem.ClientIp = transaction.Get("client_ip").String()
			//demo := transaction.Get("time_stamp").String()
			//d := time.Parse()
			datestr := transaction.Get("time_stamp").String()
			//now := time.Now().UTC().Format(time.ANSIC)
			//fmt.Println(now)
			t, err := time.Parse(TIME_LAYOUT, datestr)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(t.Format(C_TIME_LAYLOUT))
			fmt.Println(datestr)

			//result, _ := dateutil.DateTime("VV")
			//fmt.Println(datestr)

			//fmt.Println(result)

			auditItem.ServerID = transaction.Get("server_id").String()
			auditItem.UniqId = transaction.Get("unique_id").String()
			auditItem.HostIp = transaction.Get("host_ip").String()
			auditItem.HostPort, _ = strconv.Atoi(transaction.Get("host_port").String())
			auditItem.RequestHeaders = transaction.Get("request.headers").String()
			auditItem.RequestMethod = transaction.Get("request.method").String()
			auditItem.HttpVersion = transaction.Get("request.http_version").String()
			auditItem.RequestUri = transaction.Get("request.uri").String()
			auditItem.RequestBody = transaction.Get("request.body").String()
			auditItem.ResponseBody = transaction.Get("response.body").String()
			auditItem.ResponseHeaders = transaction.Get("response.headers").String()
			auditItem.AuditMsg = transaction.Get("messages").String()
			auditItem.WAFInfo = transaction.Get("producer").String()

			return true
		})

		//fmt.Println(*auditItem)

	}
	//fmt.Println(lines[0])

}


func TestJson(){
	json := "{\"transaction\":{\"client_ip\":\"10.25.8.176\",\"time_stamp\":\"Thu Jul  1 04:03:29 2021\",\"server_id\":\"065a9d91827de80ee8bc8f3fd867b63de0e6f978\",\"client_port\":33279,\"host_ip\":\"172.19\n.0.3\",\"host_port\":8080,\"unique_id\":\"1625112209\",\"request\":{\"method\":\"GET\",\"http_version\":1.1,\"uri\":\"/DVWA/vulnerabilities/sqli/?id=1%271%3D1+&Submit=Submit\",\"body\":\"\",\"header\ns\":{\"User-Agent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36\",\"Sec-Fetch-Site\":\"same-origin\",\"Upgrade-\nInsecure-Requests\":\"1\",\"sec-ch-ua-mobile\":\"?0\",\"Accept\":\"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-e\nxchange;v=b3;q=0.9\",\"X-Forwarded-Scheme\":\"https\",\"sec-ch-ua\":\"\\\" Not;A Brand\\\";v=\\\"99\\\", \\\"Google Chrome\\\";v=\\\"91\\\", \\\"Chromium\\\";v=\\\"91\\\"\",\"Sec-Fetch-User\":\"?1\",\"Connection\"\n:\"upgrade\",\"Sec-Fetch-Mode\":\"navigate\",\"X-Forwarded-For\":\"192.168.10.1\",\"X-Real-IP\":\"192.168.10.1\",\"Host\":\"192.168.10.8\",\"Sec-Fetch-Dest\":\"document\",\"Referer\":\"https://192.16\n8.10.8/DVWA/vulnerabilities/sqli/\",\"Accept-Encoding\":\"gzip, deflate, br\",\"Cookie\":\"security=low; PHPSESSID=eerefkccs2vijgcmchooiqffl0\",\"Accept-Language\":\"zh-CN,zh;q=0.9\"}},\"r\nesponse\":{\"body\":\"<html>\\r\\n<head><title>302 Found</title></head>\\r\\n<body>\\r\\n<center><h1>302 Found</h1></center>\\r\\n<hr><center>nginx</center>\\r\\n</body>\\r\\n</html>\\r\\n\",\"h\nttp_code\":302,\"headers\":{\"Server\":\"nginx\",\"Date\":\"Thu, 01 Jul 2021 04:03:29 GMT\",\"Content-Length\":\"138\",\"Content-Type\":\"text/html\",\"Connection\":\"keep-alive\",\"Location\":\"/inte\nrcepter.html\",\"Access-Control-Allow-Headers\":\"X-Requested-With\"}},\"producer\":{\"modsecurity\":\"ModSecurity v3.0.4 (Linux)\",\"connector\":\"ModSecurity-nginx v1.0.1\",\"secrules_engi\nne\":\"Enabled\",\"components\":[\"OWASP_CRS/3.3.0\\\"\"]},\"messages\":[{\"message\":\"\",\"details\":{\"match\":\"Matched \\\"Operator `Rx' with parameter `(?i:[\\\\s'\\\\\\\"`()]*?\\\\b([\\\\d\\\\w]+)\\\\b[\\\n\\s'\\\\\\\"`()]*?(?:<(?:=(?:[\\\\s'\\\\\\\"`()]*?(?!\\\\b\\\\1\\\\b)[\\\\d\\\\w]+|>[\\\\s'\\\\\\\"`()]*?(?:\\\\b\\\\1\\\\b))|>?[\\\\s'\\\\\\\"`()]*?(?!\\\\b\\\\1\\\\b)[\\\\d\\\\w]+)|(?:not\\\\s+(?:regexp|like)|is\\\\s+not|>=?|\n!=|\\\\^)[\\\\s'\\\\\\\"`()]*?(?!\\\\ (78 characters omitted)' against variable `ARGS:id' (Value: `1'1=1 ' )\",\"reference\":\"o1,4o2,1v35,6\",\"ruleId\":\"942130\",\"file\":\"conf/modsec/rules/RE\nQUEST-942-APPLICATION-ATTACK-SQLI.conf\",\"lineNumber\":\"570\",\"data\":\"\",\"severity\":\"0\",\"ver\":\"OWASP_CRS/3.3.0\",\"rev\":\"\",\"tags\":[],\"maturity\":\"0\",\"accuracy\":\"0\"}}]}}\n"
	_data := gjson.Get(json, "transaction.client_ip")
	fmt.Println(_data)
}
