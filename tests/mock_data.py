import os
from datetime import datetime
from uuid import uuid4
import time
import json
import random


# TODO 随机产生 `modsecurity` 的审计日志
def genarate_modsec_audit_item():
    data = {"transaction":{"client_ip":"10.25.8.188","time_stamp":"Mon Apr 26 08:10:25 2021","server_id":"065a9d91827de80ee8bc8f3fd867b63de0e6f978","client_port":57572,"host_ip":"172.19.0.3","host_port":8080,"unique_id":"1619424625","request":{"method":"GET","http_version":1.1,"uri":"/DVWA/vulnerabilities/fi/?page=../../e","body":"","headers":{"Host":"10.27.106.173:8080","Connection":"keep-alive","Upgrade-Insecure-Requests":"1","User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36","Cookie":"security=low; PHPSESSID=imvhgppfgj8css4aspf49sau74","Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9","Accept-Encoding":"gzip, deflate","Accept-Language":"zh-CN,zh;q=0.9"}},"response":{"body":"<html>\r\n<head><title>302 Found</title></head>\r\n<body>\r\n<center><h1>302 Found</h1></center>\r\n<hr><center>nginx</center>\r\n</body>\r\n</html>\r\n","http_code":302,"headers":{"Server":"nginx","Date":"Mon, 26 Apr 2021 08:10:25 GMT","Content-Length":"138","Content-Type":"text/html","Connection":"keep-alive","Location":"/intercepter.html","Access-Control-Allow-Headers":"X-Requested-With"}},"producer":{"modsecurity":"ModSecurity v3.0.4 (Linux)","connector":"ModSecurity-nginx v1.0.1","secrules_engine":"Enabled","components":["OWASP_CRS/3.3.0\""]},"messages":[{"message":"Path Traversal Attack (/../)","details":{"match":"Matched \"Operator `Rx' with parameter `(?i)(?:\\x5c|(?:%(?:c(?:0%(?:[2aq]f|5c|9v)|1%(?:[19p]c|8s|af))|2(?:5(?:c(?:0%25af|1%259c)|2f|5c)|%46|f)|(?:(?:f(?:8%8)?0%8|e)0%80%a|bg%q)f|%3(?:2(?:%(?:%6|4)6|F)|5%%63)|u(?:221[56]|002f|EFC8|F025)|1u|5 (400 characters omitted)' against variable `ARGS:page' (Value: `../../e' )","reference":"o33,4v4,38o2,4v35,7","ruleId":"930100","file":"conf/modsec/rules/REQUEST-930-APPLICATION-ATTACK-LFI.conf","lineNumber":"29","data":"Matched Data: /../ found within ARGS:page: ../../e","severity":"2","ver":"OWASP_CRS/3.3.0","rev":"","tags":["application-multi","language-multi","platform-multi","attack-lfi","paranoia-level/1","OWASP_CRS","capec/1000/255/153/126"],"maturity":"0","accuracy":"0"}}]}}
    data["transaction"]["client_ip"] = "1.1.1.1"
    data["transaction"]["client_port"] = 33
    _now = datetime.now()
    data["transaction"]["time_stamp"] = _now.strftime("%a ") + _now.strftime("%B")[:3] + _now.strftime(" %d %H:%M:%S %Y")
    data["transaction"]["unique_id"] = str(uuid4())
    data["transaction"]["request"]["uri"] = "/demo_url_mock"
    return data


def write_data_tofile(rand):
    with open('D:\\home\\audit_modsec.log', 'a') as f:
        for i in range(rand):
            f.write(json.dumps(genarate_modsec_audit_item()))
            f.write('\n')
    f.close()


def test():
    while True:
        print('---------write--------')
        rand = random.random()
        write_data_tofile(int(10 * rand))
        time.sleep(20 * rand)


if __name__ == '__main__':
    test()

