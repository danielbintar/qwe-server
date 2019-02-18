package region

import (
	"io/ioutil"
	"sync"

	"github.com/danielbintar/qwe-server/model"

	"gopkg.in/yaml.v2"
)

var (
	config *RegionConfig
	once sync.Once
)

type RegionConfig struct {
	Regions []*model.Region `yaml:"regions"`
}

func Instance() *RegionConfig {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("config/region/region.yaml")
		if err != nil { panic(err) }
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil { panic(err) }
	})

	return config
}
