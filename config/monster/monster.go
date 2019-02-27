package monster

import (
	"io/ioutil"
	"sync"

	"github.com/danielbintar/qwe-server/model"

	"gopkg.in/yaml.v2"
)

var (
	config *MonsterConfig
	once sync.Once
)

type MonsterConfig struct {
	Monsters []*model.Monster `yaml:"monsters"`
}

func Instance() *MonsterConfig {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("config/monster/monster.yaml")
		if err != nil { panic(err) }
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil { panic(err) }
	})

	return config
}
