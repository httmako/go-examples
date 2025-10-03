package handler

import (
    "fmt"
    "os"
	"strings"
)

type Function struct {
    Cmd string
    Help string
    Func func([]string)
}

var funcs = map[string]Function{}
func Register(cmd string, help string, f func([]string)) {
    funcs[cmd] = Function{
        Cmd: cmd,
        Help: help,
        Func: f,
    }
}

func ExecuteCommand() {
    if len(os.Args) < 2 {
        fmt.Println("available commands:\n")
		for _,v := range funcs {
			fmt.Printf("%s\n\t%s\n",v.Cmd,v.Help)
		}
		fmt.Println("")
		Exit("please provide a command")
    }
    if _, ok := funcs[os.Args[1]]; !ok {
        Exit("command not found")
        return
    }
    funcs[os.Args[1]].Func(os.Args[2:])
}

func Exit(msg string) {
    fmt.Printf("error: %s\n",msg)
    os.Exit(1)
}

func Must(err error) {
    if err != nil { panic(err) }
}

//util
func GetOnlyOneOrExit(input []string) string {
	if len(input) != 1 {
		Exit("only one argument allowed")
	}
	return input[0]
}

func GetFirstAndTailOrExit(input []string) (string,string) {
	if len(input) < 1 {
		Exit("not enough arguments, need atleast 1")
	}
	first := input[0]
	tail := ""
	if len(input) > 1 {
		tail = strings.Join(input[1:]," ")
	}
	return first, tail
}
