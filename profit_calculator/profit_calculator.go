package main

import "fmt"

const inflationRate = 2.5

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Printf("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Printf("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Printf("Tax Rate: ")
	fmt.Scan(&taxRate)

	profit := revenue - expenses
	tax := profit * (taxRate) / 100
	realProfit := profit - tax

	fmt.Printf("Profit: %d\n", realProfit)

}
