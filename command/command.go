package command

import (
	"flag"
	"fmt"
)

var commandList []Command

type Command struct {
	Name    string
	Handler CommandHandler
}

type CommandHandler func()

func Handler(name string, handle CommandHandler) {
	commandList = append(commandList, Command{
		Name: name, Handler: handle,
	})
}

func Run() {

	output := flag.String("d", "", "informasi yang akan di tampilkan, misalnya d=os, d=memory")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	if o := *output; o != "" {
		for _, v := range commandList {
			if o == v.Name {
				v.Handler()
				return
			}
		}
		fmt.Println("command not found")
	}
}
