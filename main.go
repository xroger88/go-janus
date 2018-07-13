package main

import (
	"fmt"

	"github.com/xroger88/go-janus/config"
	"github.com/xroger88/go-janus/myflag"
)

func init() {
}

func main() {

	fmt.Println("I will make go-janus by referring janus source tree")

	if myflag.Myflags.Show_help {
		myflag.PrintHelpMessage()
	}

	if myflag.Myflags.Show_version {
		fmt.Println("Version: 0.1")
	}

	myflag.PrintAll()

	// load configuration from the config file
	config.LoadConfig(myflag.Myflags.Config_file_path)

}
