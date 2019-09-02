package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"../Common"
	"../Config"
)

var (
	typeInfoStorage      map[string][]byte
	typeStorage          []byte
	cacheRW              = new(sync.RWMutex)
	cacheRefreshDuration = time.Duration(10) * time.Minute
)

func init() {
	go cacheLoader()
}
func cacheLoader() {
	ticker := time.Tick(cacheRefreshDuration)
	for {
		loadCache()
		<-ticker
	}
}
func loadCache() {
	// GetType
	res := Common.MySql{}.GetConn().Select("hotData2", []string{"name", "id"}).QueryAll()
	typeS := Common.Message{}.Success("获取数据成功", res)
	// GetTypeInfo
	typeInfoS := make(map[string][]byte)
	data := Common.MySql{}.GetConn().ExecSql("SELECT id,str FROM hotData")
	for _, one := range data {
		typeInfoS[one["id"]] = []byte(one["str"])
	}
	// Replace
	cacheRW.Lock()
	defer cacheRW.Unlock()
	typeStorage = typeS
	typeInfoStorage = typeInfoS
}

func GetTypeInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := r.ParseForm()
	if err != nil {
		log.Fatal("系统错误" + err.Error())
	}
	id := r.Form.Get("id")

	var rsp []byte
	cacheRW.RLock()
	info, ok := typeInfoStorage[id]
	if ok {
		rsp = make([]byte, len(info))
		copy(rsp, info)
	} else {
		rsp = []byte(`{"Code":1,"Message":"id错误，无该分类数据","Data":[]}`)
	}
	cacheRW.RUnlock()
	w.Write(rsp)
}

func GetType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	cacheRW.RLock()
	rsp := make([]byte, len(typeStorage))
	copy(rsp, typeStorage)
	cacheRW.RUnlock()
	w.Write(rsp)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Fprintf(w, "%s", Config.MySql().Source)
}

/**
kill -SIGUSR1 PID 可平滑重新读取mysql配置
*/
func SyncMysqlCfg() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			Config.ReloadConfig()
			log.Println("Reloaded config")
		}
	}()
}

func main() {
	SyncMysqlCfg()
	http.HandleFunc("/GetTypeInfo", GetTypeInfo) // 设置访问的路由
	http.HandleFunc("/GetType", GetType)         // 设置访问的路由
	http.HandleFunc("/GetConfig", GetConfig)     // 设置访问的路由
	err := http.ListenAndServe(":9090", nil)     // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
