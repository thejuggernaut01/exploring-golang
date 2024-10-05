package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	
	if (err != nil) {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	
	if (err != nil) {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter tax rate (e.g., 0.1 for 10%): ")

	if (err != nil) {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	os.WriteFile("results.txt", []byte(fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f", ebt, profit, ratio)), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64)  {
	ebt := revenue - expenses
	profit := ebt * taxRate
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error)  {
	var userInput float64
	fmt.Print((infoText))
	fmt.Scan((&userInput))
	
	if (userInput <= 0) {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}