package Config

type Redis struct {
}

func (redis Redis) GetConfig() map[interface{}]interface{} {
	redisCfg := make(map[interface{}]interface{})
	redisCfg["host"] = "ip"
	redisCfg["port"] = "port"
	redisCfg["size"] = 20
	return redisCfg
}
