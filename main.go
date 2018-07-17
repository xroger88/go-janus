package main

import (
	"fmt"

	"github.com/xroger88/go-janus/cmdflag"
	"github.com/xroger88/go-janus/config"
)

func init() {
}

func main() {

	fmt.Println("*** I will make go-janus by referring janus source tree ***")

	// check simple flags
	if cmdflag.Flags.Show_help {
		cmdflag.PrintHelpMessage()
		return
	}

	if cmdflag.Flags.Show_version {
		fmt.Println("Version: 0.1 (2018-07-17)")
		return
	}

	cmdflag.PrintAll()

	// load configuration from the config file
	config.LoadConfig(cmdflag.Flags.Config_file)

	// override command-line flags into configuration
	config.Conf.General.Daemonize = cmdflag.Flags.Enable_daemon

	// show the configuration set overrided by command-line flags
	config.PrintAll()

}
