package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golevi/cfallow/cmds"
)

func main() {
	if len(os.Args) < 2 {
		cmds.AddMyIP()
		return
	}

	switch os.Args[1] {
	case "myip":
		cmds.AddMyIP()
	case "file":
		fileCmd := flag.NewFlagSet("file", flag.ExitOnError)
		fileName := fileCmd.String("name", "", "filename")
		fileCmd.Parse(os.Args[2:])
		if *fileName == "" {
			fmt.Println("Expected -name argument")
			os.Exit(1)
		}
		cmds.AddFile(*fileName)

	default:
		fmt.Println("Expected 'myip' or 'file'")
		os.Exit(1)
	}
}
