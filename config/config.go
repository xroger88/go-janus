package config

// configuration must be easy for user ...
// Janus's configuration file is a little bit complex so need to improve ...
// In there, a set of section is managed, for example, general, plugins, events, transports etc.
// One of ways to configure those of sections seems YAML based structure which is easy to handle.
// Let's change the current Janus cfg file to YAML file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
	Box  struct {
		A     int   `yaml:"a"`
		Blist []int `yaml:"blist"`
	} `yaml:"box"`
}

func (c *conf) getConf(cfp string) *conf {

	yamlFile, err := ioutil.ReadFile(cfp)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *conf) setConf(cfp string) *conf {
	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	err = ioutil.WriteFile(cfp, d, os.ModePerm)
	if err != nil {
		log.Printf("yamlFile.Set err   #%v ", err)
	}

	return c
}

var c conf

func LoadConfig(filepath string) {
	c.getConf(filepath)
	fmt.Println(c)
}

func SaveConfig(filepath string) {
	//c.Time = 12345
	c.setConf(filepath)
}
