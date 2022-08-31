package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	// "os"
	"flag"
	"time"
)

var host = flag.String("host", "", "host参数")
var url = flag.String("url", "none", "url链接(不带域名时，需填写host参数)")
var thread = flag.Int64("p", 8, "并发线程数")

var speed_download int64

func main() {
	flag.Parse()
	for t := 0; t < 1; t++ {
		go download()
	}

	time.Sleep(time.Duration(30000) * time.Minute)

}

func download() {
	for true {
		client := &http.Client{}
		req, err := http.NewRequest("GET", *url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("User-Agent", "GolangSpider/1.0")
		req.Host = *host
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Add("Accept-Charset", "utf-8")
		req.Header.Add("Accept-Language", "zh-cn")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
		resp, err := client.Do(req)
		if err != nil {

		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}

}
