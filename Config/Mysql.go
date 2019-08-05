package Config

type Mysql struct {
}

func (mysql Mysql) GetConfig() map[interface{}]interface{} {
	config := make(map[interface{}]interface{})
	config["source"] = "root:password@tcp(ip:port)/database?charset=utf8mb4"
	config["driver"] = "mysql"
	return config
}
