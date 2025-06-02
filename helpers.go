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

func Update(args []string) {
	addCmd := flag.NewFlagSet("update", flag.ExitOnError)

	id := addCmd.Int("id", -1, "Id of the expense")
	amount := addCmd.Float64("amount", -1, "Amount of the expense")
	desc := addCmd.String("description", "", "Description of the Expense")
	addCmd.Parse(args)
	if *id == -1 {
		fmt.Println("Error: --id is required")
		addCmd.Usage()
		os.Exit(1)
	}
	if *amount == -1 && *desc == "" {
		fmt.Println("Error: No flag provided with update")
		addCmd.Usage()
		os.Exit(1)
	}

	var err error = nil
	if *amount == -1 {
		err = UpdateExpense(*id, *desc)
	} else {
		err = UpdateExpense(*id, *desc, *amount)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Expense Updated successfully (ID: %v)\n", *id)
}

func Delete(args []string) {
	addCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	id := addCmd.Int("id", -1, "Id of the expense")
	addCmd.Parse(args)
	if *id == -1 {
		fmt.Println("Error: --id is required")
		addCmd.Usage()
		os.Exit(1)
	}

	err := DeleteExpense(*id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Expense deleted successfully (ID: %v)\n", *id)
}

func List(args []string) {
	expenses := LoadJson()
	PrintMajorseparator()
	fmt.Println("List of Expenses")
	PrintMajorseparator()
	for _, v := range expenses.List {
		fmt.Println()
		fmt.Printf("ID: %v\n", v.Id)
		fmt.Printf("Date: %v\n", v.Date)
		fmt.Printf("Description: %v\n", v.Description)
		fmt.Printf("Amount: %v\n", v.Amount)
		fmt.Println()
		PrintMinorseparator()
	}
}

func Summary(args []string) {
	addCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	month := addCmd.Int("month", -1, "Month of the expense")
	addCmd.Parse(args)
	total := 0
	if *month == -1 {
		total = ExpenseSummary()
		fmt.Printf("Total expenses: %v\n", total)
	} else {
		total = ExpenseSummary(*month)
		fmt.Printf("Total expenses for %v: %v\n", MonthstoString[*month], total)
	}

}
