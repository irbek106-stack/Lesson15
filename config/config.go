package config

import "errors"

func GetPortFromConfig(config map[string]string) (string, error) {
	key := "PORT"
	port, ok := config[key]
	if ok {
		return port, nil
	} else {
		return port, errors.New("ключ 'PORT' отсутсвует в конфигурации")
	}
}