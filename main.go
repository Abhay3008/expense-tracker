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
	Amount      float32 `json:"amount"`
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

const filename = "./expense.json"

func main() {
	// err := AddExpense(23, "ciggarate")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// err = DeleteExpense(1)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	total := ExpenseSummary(4)
	fmt.Print(total)
}

func AddExpense(amount float32, description string) error {
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
	fmt.Println(expenses.List)
	err := SaveJson(expenses)
	if err != nil {
		fmt.Print(err)
	}
	return nil

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
