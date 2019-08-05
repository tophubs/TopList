package main

import (
	"../Common"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HotData struct {
	Code    int
	Message string
	Data    interface{}
}

func SaveDataToJson(data interface{}, dataType string) string {
	Message := HotData{}
	Message.Code = 0
	Message.Message = "获取成功"
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列号json错误")
	}
	return string(jsonStr)

}

// V2EX
func GetV2EX() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	document, err := goquery.NewDocumentFromReader(res.Body)
	var allData []map[string]interface{}
	document.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com/" + url})
		}
	})
	return allData
}

func GetITHome() []map[string]interface{} {
	url := "https://www.ithome.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	document, err := goquery.NewDocumentFromReader(res.Body)
	var allData []map[string]interface{}
	document.Find(".hot-list .bx ul li").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": url})
		}
	})
	return allData
}



// 贴吧
func GetTieBa() []map[string]interface{} {
	url := "http://tieba.baidu.com/hottopic/browse/topicList"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	var allData []map[string]interface{}
	i := 1
	for i < 30 {
		test := js.Get("data").Get("bang_topic").Get("topic_list").GetIndex(i).MustMap()
		allData = append(allData, map[string]interface{}{"title": test["topic_name"], "url": test["topic_url"]})
		i++
	}
	return allData

}

func GetChouTi() []map[string]interface{} {
	url := "https://dig.chouti.com/top/24hr?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "163"
	url2 := "https://dig.chouti.com/link/hot?afterTime=" + strconv.FormatInt(time.Now().Unix(), 10) + "026000" + "&_=" + strconv.FormatInt(time.Now().Unix(), 10) + "667"
	res, err := http.Get(url)
	res2, _ := http.Get(url2)
	if err != nil {
		log.Fatal(err.Error())
	}
	str, _ := ioutil.ReadAll(res.Body)
	str2, _ := ioutil.ReadAll(res2.Body)
	js, err2 := simplejson.NewJson(str)
	js2, _ := simplejson.NewJson(str2)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	var allData []map[string]interface{}
	i := 1
	for i < 30 {
		test := js.Get("data").GetIndex(i).MustMap()
		if test["title"] != nil && test["url"] != nil {
			allData = append(allData, map[string]interface{}{"title": test["title"], "url": test["url"]})
		}
		i++
	}
	j := 1
	for j < 60 {
		test := js2.Get("data").GetIndex(j).MustMap()
		if test["title"] != nil && test["url"] != nil {
			allData = append(allData, map[string]interface{}{"title": test["title"], "url": test["url"]})
		}
		j++
	}
	return allData

}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GetAllData(dataType string) {
	start := time.Now()
	switch dataType {
	case "V2EX":
		Common.MySql{}.GetConn().Where(map[string]string{"dataType": dataType}).Update("hotData", map[string]string{"str": SaveDataToJson(GetV2EX(), dataType)})
		group.Done()
		seconds := time.Since(start).Seconds()
		fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, dataType)
		fmt.Println()
		break
	case "TieBa":
		Common.MySql{}.GetConn().Where(map[string]string{"dataType": dataType}).Update("hotData", map[string]string{"str": SaveDataToJson(GetTieBa(), dataType)})
		group.Done()
		seconds := time.Since(start).Seconds()
		fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, dataType)
		fmt.Println()
		break
	case "ChouTi":
		Common.MySql{}.GetConn().Where(map[string]string{"dataType": dataType}).Update("hotData", map[string]string{"str": SaveDataToJson(GetChouTi(), dataType)})
		group.Done()
		seconds := time.Since(start).Seconds()
		fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, dataType)
		fmt.Println()
		break
	case "ITHome":
		Common.MySql{}.GetConn().Where(map[string]string{"dataType": dataType}).Update("hotData", map[string]string{"str": SaveDataToJson(GetITHome(), dataType)})
		group.Done()
		seconds := time.Since(start).Seconds()
		fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, dataType)
		fmt.Println()
		break
	}
}

func DeleteRedisCache() {
	sql := "select id,name from hotData"
	data := Common.MySql{}.GetConn().ExecSql(sql)
	redisInstance := Common.GetRedisConn()
	for _, value := range data {
		r, err := redisInstance.Do("DEL", "json"+value["id"])
		if err != nil {
			log.Println(err)
		} else {
			fmt.Print("删除缓存数据成功", value["name"], r)
		}
	}
}

var group sync.WaitGroup

func main() {
	allData := []string{"TieBa", "V2EX", "ChouTi", "ITHome",}
	fmt.Println("开始抓取" + strconv.Itoa(len(allData)) + "种数据类型")
	group.Add(len(allData))
	for _, value := range allData {
		fmt.Println("开始抓取" + value)
		go GetAllData(value)
	}
	group.Wait()
	DeleteRedisCache()
	fmt.Print("完成抓取")
}
