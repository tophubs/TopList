package Config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type MysqlCfg struct {
	Source, Driver string
}

var (
	cfg     *MysqlCfg
	once    sync.Once
	cfgLock = new(sync.RWMutex)
)

func MySql() *MysqlCfg {
	once.Do(ReloadConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

func GetMySqlFilePath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal("获取目录失败")
	}
	if strings.HasSuffix(currentPath, "App") {
		return strings.ReplaceAll(currentPath, "App", "Config") + "/mysql.toml"
	}
	return currentPath + "/Config/mysql.toml"
}

func ReloadConfig() {
	filePath, err := filepath.Abs(GetMySqlFilePath())
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse mysql config once. filePath: %s\n", filePath)
	config := new(MysqlCfg)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
}
