package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"
	"github.com/tophubs/TopList/Common"
	"github.com/tophubs/TopList/Config"
)

func GetTypeInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("系统错误" + err.Error())
	}
	id := r.Form.Get("id")
	re := regexp.MustCompile("[0-9]+")
	id = re.FindString(id)
	sql := "select str from hotData2 where id=" + id
	data := Common.MySql{}.GetConn().ExecSql(sql)
	if len(data) == 0 {
		fmt.Fprintf(w, "%s", `{"Code":1,"Message":"id错误，无该分类数据","Data":[]}`)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", data[0]["str"])
}

func GetType(w http.ResponseWriter, r *http.Request) {
	res := Common.MySql{}.GetConn().Select("hotData2", []string{"name", "id"}).QueryAll()
	Common.Message{}.Success("获取数据成功", res, w)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", Config.MySql().Source)
}

/**
kill -SIGUSR1 PID 可平滑重新读取mysql配置
*/
//func SyncMysqlCfg() {
//	s := make(chan os.Signal, 1)
//	signal.Notify(s, syscall.SIGUSR1)
//	go func() {
//		for {
//			<-s
//			Config.ReloadConfig()
//			log.Println("Reloaded config")
//		}
//	}()
//}

func main() {
	//SyncMysqlCfg()
	http.HandleFunc("/GetTypeInfo", GetTypeInfo) // 设置访问的路由
	http.HandleFunc("/GetType", GetType)         // 设置访问的路由
	http.HandleFunc("/GetConfig", GetConfig)     // 设置访问的路由

	// 静态资源
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../Html/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../Html/js/"))))

	// 首页
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("../Html/hot.html")
		if err != nil {
			log.Println("err")
		}
		t.Execute(res, nil)
	})

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
