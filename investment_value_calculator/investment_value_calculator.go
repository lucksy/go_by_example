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

	fmt.Print("eExpected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
