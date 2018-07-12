package main

import (
	"fmt"

	"github.com/xroger88/go-janus/myflag"
)

var gbHelp, gbVersion, gbDaemon bool
var gbPidFilePath string

func init() {
	myflag.BoolVar(&gbHelp, []string{"h", "help"}, false, "Print help and exit")
	myflag.BoolVar(&gbVersion, []string{"v", "version"}, false, "Print version and exit")
	myflag.BoolVar(&gbDaemon, []string{"b", "daemon"}, false, "Launch Janus in background as a dameon (default=off)")
	myflag.StringVar(&gbPidFilePath, []string{"p", "pid-file"}, "", "Open the specified PID file `path` when starting Janus")
	myflag.Parse()
}

func main() {

	fmt.Println("I will make go-janus by referring janus source tree")

	fmt.Printf("gbHelp=%v\n", gbHelp)
	fmt.Printf("gbVersion=%v\n", gbVersion)
	fmt.Printf("gbDaemon=%v\n", gbDaemon)
	fmt.Printf("gbPidFilePath=%v\n", gbPidFilePath)
}
