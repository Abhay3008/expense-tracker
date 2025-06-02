# Expense Tracker CLI

A simple command-line tool for tracking personal expenses, written in Go. Expenses are stored in a local JSON file with support for adding, updating, deleting, listing, and summarizing expenses.

## Features

- Add new expenses with description and amount  
- Update existing expenses  
- Delete expenses by ID  
- List all recorded expenses  
- View total expenses or filter by month  

## Usage

```bash
expense-tracker <command> [options]
```

## Commands

### Add an expense

```bash
expense-tracker add --description "Lunch" --amount 20
```

### Update an expense

```bash
expense-tracker update --id 2 --description "Dinner" --amount 10
```

### Delete an expense

```bash
expense-tracker delete --id 2
```

### List all expenses

```bash
expense-tracker list
```

### Show total expenses

```bash
expense-tracker summary
```

### Show total expenses for a specific month

```bash
expense-tracker summary --month 8
```

### Help

```bash
expense-tracker help
```

## Expense Storage

All expenses are saved in a local `expenses.json` file in the same directory.

## Build

### To build the binary

```bash
go build -o expense-tracker
```

### To build and run with Docker

```bash
docker build -t expense-tracker .
docker run --rm -v $(pwd)/expenses.json:/app/expenses.json expense-tracker list
```

## Example

```bash
$ expense-tracker add --description "Lunch" --amount 20
Expense added successfully (ID: 1)

$ expense-tracker list
```

## Project URL

[https://roadmap.sh/projects/expense-tracker](https://roadmap.sh/projects/expense-tracker)
