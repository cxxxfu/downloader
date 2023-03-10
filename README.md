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
        nohup ./cdn -host lg-lax.fdcservers.net     -url http://192.240.120.250/100MBtest.zip -p 1
