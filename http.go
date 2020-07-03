package godash

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type objectJSON = map[string]interface{}
type transportLite struct {
	Port     string
	Host     string
	Protocol string
}

var proxyCache map[string]string
var proxyRetry = 0

func Post(path string, data map[string]string, headers map[string]string) (string, error) {

	return request("POST", path, data, headers)

}
func Get(path string, data map[string]string, headers map[string]string) (string, error) {
	query := url.URL{}
	qq := query.Query()
	for k, v := range data {
		qq.Set(k, v)
	}
	flag := strings.Index(path, "?")
	if flag < 0 {
		path = path + "?" + qq.Encode()
	} else {
		path = path + "&" + qq.Encode()
	}
	return request("GET", path, nil, headers)
}
func proxy() (*http.Transport, *transportLite) {
	var proxyPort string
	var proxyHost string
	var proxyProtocol string
	var transport *http.Transport
	//
	log.Println("cache:", proxyCache)
	if nil == proxyCache {
		// 初始化缓存
		proxyCache = make(map[string]string)
		proxyCache["port"] = "3128"
		proxyCache["host"] = "5.189.188.95"
		proxyCache["protocol"] = "http"

		proxyPort = proxyCache["port"]
		proxyHost = proxyCache["host"]
		proxyProtocol = proxyCache["protocol"]
	}

	if nil != proxyCache && "" != proxyCache["port"] {
		// 使用缓存
		proxyPort = proxyCache["port"]
		proxyHost = proxyCache["host"]
		proxyProtocol = proxyCache["protocol"]

	} else {
		// 更新缓存
		proxyResp, err := http.Get("https://ip.jiangxianli.com/api/proxy_ip")
		if err != nil {
			log.Println("请求代理失败 ", err)
			return nil, nil
		}
		proxyRet, err := ioutil.ReadAll(proxyResp.Body)
		log.Println(proxyRet)

		proxyData := make(objectJSON)
		json.Unmarshal(proxyRet, &proxyData)
		proxyItem := proxyData["data"].(objectJSON)
		proxyPort = proxyItem["port"].(string)
		proxyHost = proxyItem["ip"].(string)
		proxyProtocol = proxyItem["protocol"].(string)
		if "" == proxyProtocol {
			proxyProtocol = "http"
		}
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyProtocol + "://" + proxyHost + ":" + proxyPort)
	}
	transport = &http.Transport{Proxy: proxy}
	log.Println("代理地址: ", proxyProtocol+"://"+proxyHost+":"+proxyPort)
	return transport, &transportLite{
		Port:     proxyPort,
		Host:     proxyHost,
		Protocol: proxyProtocol,
	}

}
func request(method string, uri string, data map[string]string, headers map[string]string) (body string, err error) {
	param := make(url.Values)
	for k, v := range data {
		param[k] = []string{v}
	}
	if nil == headers {
		headers = map[string]string{
			"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36",
			"content-type": "application/x-www-form-urlencoded",
		}
	}

	log.Println(uri)
	req, err := http.NewRequest(method, uri, strings.NewReader(param.Encode()))
	if nil != headers {
		for k, v := range headers {
			req.Header.Set(k, v)

		}
	}
	//proxyResp, err := http.Get("https://ip.jiangxianli.com/api/proxy_ips")
	//proxyRet, err := ioutil.ReadAll(proxyResp.Body)
	//
	//proxyData := make(objectJSON)
	//json.Unmarshal(proxyRet, &proxyData)
	//proxyList := proxyData["data"].(objectJSON)["data"].([]interface{})
	//rand.Seed(time.Now().Unix())
	//proxyRand := rand.Int() % len(proxyList)
	//log.Println("proxyRand", proxyRand)
	//proxyItem := proxyList[proxyRand].(objectJSON)
	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse("http://" + proxyItem["ip"].(string) + ":" + proxyItem["port"].(string))
	//}

	transport, lite := proxy()
	client := &http.Client{Transport: transport, Timeout: time.Second * 10}

	//var transport *http.Transport
	//var lite *transportLite
	//client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		proxyRetry++
		// 超过5次，重新获取有效IP
		if nil != transport && proxyRetry > 5 {
			proxyCache["port"] = ""
			proxyCache["host"] = ""
			proxyCache["protocol"] = ""
			proxyRetry = 0
		}
		log.Println("Retry ...", proxyRetry)
		log.Println("Wait ...", proxyRetry*1000)
		time.Sleep(time.Duration(proxyRetry*1000) * time.Millisecond)
		return request(method, uri, data, headers)
	}
	defer resp.Body.Close()

	if nil != transport {
		proxyCache["port"] = lite.Port
		proxyCache["host"] = lite.Host
		proxyCache["protocol"] = lite.Protocol
		proxyRetry = 0
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	return string(ret), nil
}
