package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/xroger88/go-janus/cmdflag"
	"github.com/xroger88/go-janus/config"
	"github.com/xroger88/go-janus/util"
)

func init() {
}

func main() {

	// Command & Flags and Configuration Part
	// Current implementation is workable so I will keep this in some time
	// But these of stuffs will be rewriten with Cobra and Viper

	// check simple flags
	if cmdflag.Flags.Show_help {
		cmdflag.PrintHelpMessage()
		return
	}

	if cmdflag.Flags.Show_version {
		fmt.Println("Version: 0.1 (2018-07-17)")
		return
	}

	if cmdflag.Flags.Show_flags {
		cmdflag.PrintAll()
		return
	}

	// load configuration from the config file
	config.LoadConfig(cmdflag.Flags.Config_file)

	// override command-line flags into configuration
	config.Conf.General.Daemonize = cmdflag.Flags.Enable_daemon
	if cmdflag.Flags.Log_file != cmdflag.DEF_LOG_FILE {
		config.Conf.General.Log_to_file = cmdflag.Flags.Log_file
	}

	log_file := config.Conf.General.Log_to_file
	if log_file != "" {
		// init log to store logs into the specified log file
		util.LogInit(log_file)
	}

	// show the configuration set overrided by command-line flags
	if cmdflag.Flags.Show_config {
		config.PrintAll()
		return
	}

	log.Infoln("*** I will make go-janus by referring janus source tree ***")

}
