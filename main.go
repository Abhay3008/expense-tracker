package main

import (
	"fmt"
	"os"
)

const filename = "./expense.json"

func main() {

	args := os.Args[1:]
	// fmt.Println(args)
	if len(args) < 1 {
		Help()
		os.Exit(0)
	}
	switch args[0] {

	case "add":
		Add(args[1:])
	case "update":
		Update(args[1:])
	case "delete":
		Delete(args[1:])
	case "list":
		List(args[1:])
	case "summary":
		Summary(args[1:])
	case "help":
		Help()
	default:
		fmt.Printf("Invalid option: %v\n", args[0])
		Help()
	}

}
