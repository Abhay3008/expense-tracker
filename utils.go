package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Id          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

type ExpenseList struct {
	List    []Expense
	TotalId int `json:"totalId"`
}

var MonthsMap = map[int]string{
	1:  "01",
	2:  "02",
	3:  "03",
	4:  "04",
	5:  "05",
	6:  "06",
	7:  "07",
	8:  "08",
	9:  "09",
	10: "10",
	11: "11",
	12: "12",
}
var MonthstoString = map[int]string{
	1:  "Jan",
	2:  "Feb",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "Aug",
	9:  "Sept",
	10: "Oct",
	11: "Nov",
	12: "Dec",
}

func AddExpense(amount float64, description string) (int, error) {
	expenses := LoadJson()
	date := time.Now().Format("2006-01-02")
	var expense = Expense{
		Id:          expenses.TotalId + 1,
		Amount:      amount,
		Description: description,
		Date:        date,
	}
	expenses.List = append(expenses.List, expense)
	expenses.TotalId++
	id := expense.Id
	err := SaveJson(expenses)
	if err != nil {
		fmt.Print(err)
	}
	return id, nil

}

func UpdateExpense(id int, desc string, amount ...float64) error {
	expenses := LoadJson()
	for i, v := range expenses.List {
		if v.Id == id {
			if len(amount) > 0 {
				expenses.List[i].Amount = amount[0]
			}
			if desc != "" {
				expenses.List[i].Description = desc
			}
			err := SaveJson(expenses)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			return nil
		}
	}
	return errors.New("unable to find expense with id")
}

func DeleteExpense(id int) error {
	expenses := LoadJson()
	for i, v := range expenses.List {
		if v.Id == id {
			expenses.List = append(expenses.List[0:i], expenses.List[i+1:]...)
			err := SaveJson(expenses)
			if err != nil {
				fmt.Print(err)
			}
			return nil
		}
	}
	return errors.New("unable to find expense with id")
}

func ExpenseSummary(month ...int) int {
	expenses := LoadJson()
	total := 0
	if len(month) < 1 {
		for _, v := range expenses.List {
			total += int(v.Amount)
		}
	} else {
		year := strconv.Itoa(time.Now().Year())
		month := MonthsMap[month[0]]
		dateprefix := year + "-" + month
		for _, v := range expenses.List {
			if strings.HasPrefix(v.Date, dateprefix) {
				total += int(v.Amount)
			}
		}
	}
	return total

}

func LoadJson() ExpenseList {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return ExpenseList{
			List:    []Expense{},
			TotalId: 0,
		}
	} else {
		readfile, _ := os.ReadFile(filename)
		var list ExpenseList
		json.Unmarshal(readfile, &list)
		return list
	}
}

func SaveJson(expenselist ExpenseList) error {
	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	res, err := json.MarshalIndent(expenselist, "", "\t")
	if err != nil {
		return err
	}
	io.Writer.Write(newfile, res)
	newfile.Close()
	return nil
}

func Help() {
	var text = `Usage: expense-tracker [command] [options]

Commands:
  add       Add a new expense
            --description <text>   Description of the expense (required)
            --amount <number>      Amount of the expense (required)

  update    Update an existing expense
            --id <number>          ID of the expense to update (required)
            --description <text>   New description (optional)
            --amount <number>      New amount (optional)

  delete    Delete an expense
            --id <number>          ID of the expense to delete (required)

  list      List all expenses

  summary   Show total expenses
            --month <number>       (optional) Filter total by month (1–12)

  help      Show this help message

Examples:
  expense-tracker add --description "Lunch" --amount 20
  expense-tracker update --id 2 --description "Dinner" --amount 10
  expense-tracker delete --id 2
  expense-tracker list
  expense-tracker summary
  expense-tracker summary --month 8
`
	fmt.Println(text)
}

func PrintMajorseparator() {
	fmt.Println("---------------")
}

func PrintMinorseparator() {
	fmt.Println("-------")
}
