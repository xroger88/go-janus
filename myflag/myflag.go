package myflag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]...\n", os.Args[0])
		PrintDefaults()
	}
}

// BoolVar defines a bool flag with specified mulitple names, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, names []string, value bool, usage string) {
	for _, name := range names {
		flag.CommandLine.BoolVar(p, name, value, usage)
	}
}

// StringVar defines a string flag with specified multiple names, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, names []string, value string, usage string) {
	for _, name := range names {
		flag.CommandLine.StringVar(p, name, value, usage)
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

		/**
		if !flag.isZeroValue(f, f.DefValue) {
			if _, ok := f.Value.(*stringValue); ok {
				// put quotes on the value
				s += fmt.Sprintf(" (default %q)", f.DefValue)
			} else {
				s += fmt.Sprintf(" (default %v)", f.DefValue)
			}
		}
		**/
		fmt.Fprint(flag.CommandLine.Output(), s, "\n")
	})
}

// Parse parses the command-line flags from os.Args[1:]. Must be called
func Parse() {
	flag.Parse()
}
