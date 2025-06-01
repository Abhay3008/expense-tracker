package main

import (
	"flag"
	"fmt"
	"os"
)

func Add(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	desc := addCmd.String("description", "", "Description of the expense")
	amount := addCmd.Float64("amount", 0, "Amount of the expense")
	addCmd.Parse(args)
	if *desc == "" {
		fmt.Println("Error: --description is required")
		addCmd.Usage()
		os.Exit(1)
	}
	if *amount == 0 {
		fmt.Println("Error: --amount is required and must be non-zero")
		addCmd.Usage()
		os.Exit(1)
	}

	id, err := AddExpense(*amount, *desc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Expense added successfully (ID: %v)\n", id)
}
