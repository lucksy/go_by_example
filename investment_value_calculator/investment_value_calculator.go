package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var expectedReturnRate float64
	var years float64
	var investmentAmount float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Value (after inflation) %.2f\n", futureRealValue)

}
