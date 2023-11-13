package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {

	x := getInput()

	fmt.Println("Generating the series for first:", x, " numbers")

	series := fibonacci(x)

	fmt.Println("The generated series for first", x, " numbers is ", series)

}

/*
@param limit - the limit of the fibonacci series
*/
func fibonacci(limit int) []*big.Int {

	if limit <= 0 {
		return []*big.Int{}
	} else if limit == 1 {
		return []*big.Int{big.NewInt(0)}
	}

	generatedSeries := make([]*big.Int, 2, limit)

	generatedSeries[0], generatedSeries[1] = big.NewInt(0), big.NewInt(1)

	//since the first two elements of the series are already known, we calculate from the third element

	for i := 2; i < limit; i++ {
		// Add the sum of the two preceding numbers to the series

		nextNumber := new(big.Int).Add(generatedSeries[i-1], generatedSeries[i-2])
		generatedSeries = append(generatedSeries, nextNumber)
	}

	return generatedSeries

}

func getInput() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer value: ")
		scanner.Scan()
		input := scanner.Text()

		// Try to convert the input to an integer
		if value, err := strconv.Atoi(input); err == nil {
			return value
		} else {
			fmt.Println("Invalid input. Please enter a valid integer.")
		}
	}
}
