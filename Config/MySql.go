package Config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
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

func ReloadConfig() {
	filePath, err := filepath.Abs("./Config/mysql.toml")
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
