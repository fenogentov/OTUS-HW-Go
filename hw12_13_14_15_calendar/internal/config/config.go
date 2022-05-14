package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config ...
// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger     LoggerConf
	HTTPServer HTTPServerConf
	DB         DBConf
}

// LoggerConf ...
type LoggerConf struct {
	Level, File string
}

// HTTPServerConf ...
type HTTPServerConf struct {
	Host, Port string
}

// DBConf ...
type DBConf struct {
	Enable   bool
	User     string
	Password string
	Host     string
	Port     string
	NameDB   string
}

// NewConfig parsing config file.
func NewConfig(path string) (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return Config{}, err
	}
	fmt.Printf("config: %+v\n", conf)
	return conf, nil
}

// TODO
