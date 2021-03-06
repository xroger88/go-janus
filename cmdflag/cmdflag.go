package cmdflag

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/xroger88/go-janus/util"
)

var Flags struct {
	Show_help, Show_version, Enable_daemon, Show_flags, Show_config, Disable_stdout bool
	Pid_file, Log_file, Config_file                                                 string
	Http_port                                                                       uint
}

const (
	DEF_PID_FILE    = "./janus.pid"
	DEF_LOG_FILE    = "./janus.log"
	DEF_CONFIG_FILE = "./conf.yaml"
	DEF_HTTP_PORT   = 8080
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]...\n", os.Args[0])
		PrintDefaults()
	}
	BoolVar(&Flags.Show_help, []string{"h", "help"}, false, "Print help and exit")
	BoolVar(&Flags.Show_version, []string{"v", "version"}, false, "Print version and exit")
	BoolVar(&Flags.Enable_daemon, []string{"d", "daemon"}, false, "Launch Janus in background as a dameon")
	BoolVar(&Flags.Show_flags, []string{"f", "flags"}, false, "Print command line flags and exit")
	BoolVar(&Flags.Show_config, []string{"sc", "showconfig"}, false, "Print current configuration and exit")
	BoolVar(&Flags.Disable_stdout, []string{"N", "disable-stdout"}, false, "Disable stdout based logging")
	StringVar(&Flags.Pid_file, []string{"p", "pid-file"}, DEF_PID_FILE,
		"Open the specified PID file `path` when starting Janus")
	StringVar(&Flags.Log_file, []string{"l", "log-file"}, DEF_LOG_FILE,
		"Open the specified log file `path` when starting Janus")
	StringVar(&Flags.Config_file, []string{"c", "config-file"}, DEF_CONFIG_FILE,
		"Open the specified config file `path` when starting Janus")
	UintVar(&Flags.Http_port, []string{"hp", "http_port"}, DEF_HTTP_PORT,
		"Web server will be listen to http port")

	// parse the flags in command line
	flag.Parse()
}

// BoolVar defines a bool flag with specified mulitple names, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, names []string, value bool, usage string) {
	for _, name := range names {
		flag.CommandLine.BoolVar(p, name, value, usage)
	}
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, names []string, value int, usage string) {
	for _, name := range names {
		flag.CommandLine.IntVar(p, name, value, usage)
	}
}

// Int64Var defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func Int64Var(p *int64, names []string, value int64, usage string) {
	for _, name := range names {
		flag.CommandLine.Int64Var(p, name, value, usage)
	}
}

// UintVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func UintVar(p *uint, names []string, value uint, usage string) {
	for _, name := range names {
		flag.CommandLine.UintVar(p, name, value, usage)
	}
}

// Uint64Var defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func Uint64Var(p *uint64, names []string, value uint64, usage string) {
	for _, name := range names {
		flag.CommandLine.Uint64Var(p, name, value, usage)
	}
}

// StringVar defines a string flag with specified multiple names, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, names []string, value string, usage string) {
	for _, name := range names {
		flag.CommandLine.StringVar(p, name, value, usage)
	}
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, names []string, value float64, usage string) {
	for _, name := range names {
		flag.CommandLine.Float64Var(p, name, value, usage)
	}
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func DurationVar(p *time.Duration, names []string, value time.Duration, usage string) {
	for _, name := range names {
		flag.CommandLine.DurationVar(p, name, value, usage)
	}
}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// command-line flags.
// For an integer valued flag x, the default output has the form
//	-x int
//		usage-message-for-x (default 7)
// For an integer valued multiple flags but as one value such as x or xx, the default output has the form
//	-x, -xx int
//		usage-message-for-x-xx (default 7)
// The usage message will appear on a separate line for anything but
// a bool flag with a one-byte name. For bool flags, the type is
// omitted and if the flag name is one byte the usage message appears
// on the same line. The parenthetical default is omitted if the
// default is the zero value for the type. The listed type, here int,
// can be changed by placing a back-quoted name in the flag's usage
// string; the first such item in the message is taken to be a parameter
// name to show in the message and the back quotes are stripped from
// the message when displayed. For instance, given
//	flag.String("I", "", "search `directory` for include files")
// the output will be
//	-I directory
//		search directory for include files.
func PrintDefaults() {

	var touched = make(map[string]bool)
	flag.VisitAll(func(f *flag.Flag) {
		touched[f.Name] = false
	})

	flag.VisitAll(func(f *flag.Flag) {
		if touched[f.Name] {
			// already touched so skip this flag
			return
		}
		touched[f.Name] = true
		s := fmt.Sprintf("  -%s", f.Name) // Two spaces before -; see next two comments.

		flag.VisitAll(func(f2 *flag.Flag) {
			if !touched[f2.Name] {
				if strings.Compare(f.Usage, f2.Usage) == 0 {
					s += fmt.Sprintf(", --%s", f2.Name)
					touched[f2.Name] = true
				}
			}
		})

		name, usage := flag.UnquoteUsage(f)

		if len(name) > 0 {
			s += " " + name
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}
		s += strings.Replace(usage, "\n", "\n    \t", -1)

		if vk, zero := isZeroValue(f, f.DefValue); !zero {
			if vk == reflect.String {
				// put quotes on the value
				s += fmt.Sprintf(" (default %q)", f.DefValue)
			} else {
				s += fmt.Sprintf(" (default %v)", f.DefValue)
			}
		}
		fmt.Fprint(flag.CommandLine.Output(), s, "\n")
	})
}

// isZeroValue guesses whether the string represents the zero
// value for a flag. It is not accurate but in practice works OK.
// [xroger88] slightly modified the original one to return the flag value's type kind
func isZeroValue(f *flag.Flag, value string) (reflect.Kind, bool) {
	typ := reflect.TypeOf(f.Value)
	var z reflect.Value
	var fvk reflect.Kind
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
		fvk = typ.Elem().Kind()
	} else {
		z = reflect.Zero(typ)
		fvk = typ.Kind()
	}
	if value == z.Interface().(flag.Value).String() {
		return fvk, true
	}

	switch value {
	case "false", "", "0":
		return fvk, true
	}
	return fvk, false
}

func PrintHelpMessage() {
	flag.Usage()
}

func PrintAll() {
	fmt.Printf("*** CommandLine Flags ***\n")
	util.PrintValue(0, &Flags)
}
