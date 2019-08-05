package Config

type Config interface {
	GetConfig() map[interface{}]interface{}
}
