package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ArgParser struct {
	flags []StructFlags
}

type Flags struct {
	key   string
	value string
}

type StructFlags struct {
	key         string
	name        string
	value       any
	required    bool
	description string
}

func (AP ArgParser) Init() ArgParser {
	tmp := ArgParser{}
	tmp.SetFlags("-h", "--help", nil, false, "\tShow this help message")
	return tmp
}

func (AP ArgParser) ParseFlags() []Flags {
	var flags []Flags
	var ignore_i = 0

	for i, arg := range os.Args {
		if ignore_i == i {
			continue
		}

		// check flags
		if strings.HasPrefix(arg, "-") {
			// help message
			if arg == "-h" || arg == "--help" {
				AP.PrintHelp()
			}
			if strings.HasPrefix(os.Args[i+1], "-") {
				// not a value its another flag
				continue
			}
			flags = append(flags, Flags{key: arg, value: os.Args[i+1]})
			ignore_i = i + 1
		}
	}
	// check for missing flags
	missing := AP.RequiredFlagsMissing(flags)
	if len(missing) > 0 {
		fmt.Printf("Missing the following flags: %v\n", missing)
		AP.PrintHelp()
	}
	return flags
}

func (AP ArgParser) RequiredFlagsMissing(flags []Flags) []string {
	var missing_flags []string
	for _, apflag := range AP.flags {
		if !apflag.required {
			continue
		}

		found := false
		for _, flag := range flags {
			if flag.key == apflag.key {
				found = true
			}
		}
		// check if it has a default value assigned
		if apflag.value != nil {
			found = true
		}
		// missing flag
		if !found {
			missing_flags = append(missing_flags, apflag.key)
		}
	}
	return missing_flags
}

func (AP *ArgParser) SetFlags(key string, name string, default_value any, required bool, description ...string) {
	// check if default has the correct value
	AP.flags = append((*AP).flags, StructFlags{key: key,
		name:        name,
		value:       default_value,
		required:    required,
		description: description[0]})
}

func (AP *ArgParser) Parse() {
	for _, parsedflag := range AP.ParseFlags() {
		id, match := AP.KeyMatch(parsedflag.key)
		if match {
			whatAmI := func(i interface{}) {
				switch i.(type) {
				case bool:
					v, err := strconv.ParseBool(parsedflag.value)
					if err != nil {
						println("Error parsing value")
						os.Exit(990)
					}
					(*AP).flags[id].value = v
				case int:
					v, err := strconv.ParseInt(parsedflag.value, 10, 64)
					if err != nil {
						println("Error parsing value")
						os.Exit(991)
					}
					(*AP).flags[id].value = v
				case string:
					(*AP).flags[id].value = parsedflag.value
				case float32:
					v, err := strconv.ParseFloat(parsedflag.value, 64)
					if err != nil {
						println("Error parsing value")
						os.Exit(992)
					}
					(*AP).flags[id].value = v
				default:
					println("Error parsing value : unknown value type")
					os.Exit(999)
				}
			}
			whatAmI((*AP).flags[id].value)
		}
	}
}

func (AP *ArgParser) GetFieldValue(field string) any {
	for _, item := range AP.flags {
		if item.name == "--"+field {
			// fix int assetion
			switch item.value.(type) {
			case int64:
				return AP.ConvertInt64ToInt(item.value)
			default:
				return item.value
			}
		}
	}
	return nil
}

func (AP *ArgParser) ConvertInt64ToInt(value any) int {
	return int(value.(int64))
}

func (AP ArgParser) KeyMatch(key string) (int, bool) {
	for i, flag := range AP.flags {
		if flag.name == key || flag.key == key {
			return i, true
		}
	}
	return -1, false
}

func (AP ArgParser) PrintHelp() {
	final := ""
	for _, flag := range AP.flags {
		tmp := ""
		if flag.value == nil {
			tmp = fmt.Sprintf("\t%v\t%v\t%v\n", flag.key, flag.name, flag.description)
		} else {
			tmp = fmt.Sprintf("\t%v\t%v\t%v (default:%v)\n", flag.key, flag.name, flag.description, flag.value)
		}

		final += tmp
	}
	fmt.Printf("Flags available:\n%v", final)
	os.Exit(0)
}
