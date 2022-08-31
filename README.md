# downloader
刷cdn流量脚本
Usage of ./cdn:
  -host string
        host参数
  -p uint
        并发线程数 (default 8)
  -url string
        url链接(不带域名时，需填写host参数) 
        
        例如 
        nohup ./cdn -host img.rr.tv     -url http://[2408:8738:b000:11:8000:0:b00:100]/apk/20220120/o_0332650cebba4ddc80bbba5704ae7fee.apk -p 1
