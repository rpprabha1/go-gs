package configs

import (
	"encoding/json"
	m "go-gs/models"
	"io/ioutil"
)

func LoadConfig(cfgPath string) m.Configs {
	var cfg m.Configs
	raw, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic("unable to parse config: " + err.Error())
	}
	json.Unmarshal(raw, &cfg)
	return cfg
}
