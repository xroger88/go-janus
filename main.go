package main

import (
	"fmt"

	"github.com/xroger88/go-janus/myflag"
)

var gbHelp, gbVersion, gbDaemon bool
var gbPidFilePath string
var gbPort uint
var gbRadius float64

func init() {
	myflag.BoolVar(&gbHelp, []string{"h", "help"}, false, "Print help and exit")
	myflag.BoolVar(&gbVersion, []string{"v", "version"}, false, "Print version and exit")
	myflag.BoolVar(&gbDaemon, []string{"b", "daemon"}, false, "Launch Janus in background as a dameon")
	myflag.StringVar(&gbPidFilePath, []string{"p", "pid-file"}, "./janus.pid", "Open the specified PID file `path` when starting Janus")
	myflag.UintVar(&gbPort, []string{"hp", "http_port"}, 8080, "http port for testing")
	myflag.Float64Var(&gbRadius, []string{"r", "rd", "radius"}, 30.05, "for testing float64 type flag")
	myflag.Parse()
}

func main() {

	fmt.Println("I will make go-janus by referring janus source tree")

	fmt.Printf("gbHelp=%v\n", gbHelp)
	fmt.Printf("gbVersion=%v\n", gbVersion)
	fmt.Printf("gbDaemon=%v\n", gbDaemon)
	fmt.Printf("gbPidFilePath=%q\n", gbPidFilePath)
	fmt.Printf("gbPort=%v\n", gbPort)
	fmt.Printf("gbRadius=%v\n", gbRadius)
}
