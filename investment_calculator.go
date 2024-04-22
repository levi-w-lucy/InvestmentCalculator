package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	retreiveValFromUser("Enter an Investment Amount: ", &investmentAmount)
	retreiveValFromUser("Enter Years to be invested: ", &years)
	retreiveValFromUser("Enter Expected Rate of Return: ", &expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("$%.2f", futureValue)
	fmt.Printf("\n$%.2f", futureRealValue)
}

func retreiveValFromUser(message string, variablePtr *float64) {
	fmt.Print(message)
	fmt.Scan(variablePtr)
}
