package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new entry")
	flag.IntVar(&cf.Del, "del", -1, "Delete a entry")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an entry")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle an entry")
	flag.BoolVar(&cf.List, "list", false, "List all entries")

	flag.Parse()
	return &cf
}

func (cd *CmdFlags) Execute(todos *Todos) {
	switch {
	case cd.List:
		todos.print()
	case cd.Add != "":
		todos.add(cd.Add)
	case cd.Edit != "":
		parts := strings.SplitN(cd.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format, please use the 'index:new_title' format")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cd.Toggle != -1:
		todos.toggle(cd.Toggle)
	case cd.Del != -1:
		todos.delete(cd.Del)
	default:
		fmt.Println("Invalid command")
	}
}
