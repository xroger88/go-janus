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

	"github.com/xroger88/go-janus/util"
	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Name    string
	General struct {
		Configs_folder          string
		Plugins_folder          string
		Transports_folder       string
		Events_folder           string
		Log_to_stdout           bool
		Log_to_file             string
		Daemonize               bool
		Pid_file                string
		Interface               string
		Debug_level             int
		Debug_timestamps        bool
		Debug_colors            bool
		Debug_locks             bool
		Api_secret              string
		Token_auth              bool
		Token_auth_secret       string
		Admin_secret            string
		Server_name             string
		Session_timeout         int
		Reclain_session_timeout int
		Recordings_tmp_ext      string
	}
}

func (c *ConfigType) getConf(cfp string) *ConfigType {

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

func (c *ConfigType) setConf(cfp string) *ConfigType {
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

var Conf ConfigType

func LoadConfig(filepath string) {
	Conf.getConf(filepath)
	//fmt.Println(Conf)
}

func SaveConfig(filepath string) {
	//c.Time = 12345
	Conf.setConf(filepath)
}

func PrintAll() {
	fmt.Printf("*** The Configuration Details *** \n")
	util.PrintValue(0, &Conf)
}
