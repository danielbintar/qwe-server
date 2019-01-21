package town

import (
	"io/ioutil"
	"sync"

	"github.com/danielbintar/qwe-server/model"

	"gopkg.in/yaml.v2"
)

var (
	config *TownConfig
	once sync.Once
)

type TownConfig struct {
	Towns []*model.Town `yaml:"towns"`
}

func Instance() *TownConfig {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("config/town/town.yaml")
		if err != nil { panic(err) }
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil { panic(err) }
	})

	return config
}
