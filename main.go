package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	x := getInput() // Returns an integer value entered by the user

	fmt.Println("Generating the series for :", x, " numbers")

	series := fibonacci(x) // Returns a slice of pointers to big.Int

	fmt.Println("The generated series for ", x, " numbers is ", series)
}

// fibonacci calculates the Fibonacci series up to the specified limit.
// @param limit - the limit of the Fibonacci series
// @return []*big.Int - a slice of pointers to big.Int, representing the Fibonacci series
func fibonacci(limit int) []*big.Int {
	if limit <= 0 {
		return []*big.Int{} // Return an empty slice of big.Int pointers
	} else if limit == 1 {
		return []*big.Int{big.NewInt(0)} // Return a slice with the first element of the series
	}

	generatedSeries := make([]*big.Int, 2, limit) // Initialize the slice with a capacity of 'limit'

	generatedSeries[0], generatedSeries[1] = big.NewInt(0), big.NewInt(1) // Set the first two elements

	// Calculate the series from the third element onward
	for i := 2; i < limit; i++ {
		nextNumber := new(big.Int).Add(generatedSeries[i-1], generatedSeries[i-2])
		generatedSeries = append(generatedSeries, nextNumber)
	}

	return generatedSeries // Return the complete Fibonacci series
}

// getInput prompts the user for an integer input and validates it.
// @return int - the validated integer input from the user
func getInput() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer value: ")
		scanner.Scan()
		input := scanner.Text()

		// Convert and validate the input
		if value, err := strconv.Atoi(input); err == nil {
			return value // Return the integer value
		} else {
			fmt.Println("Invalid input. Please enter a valid integer.")
		}
	}
}
