# Fibonacci Series Generator in Go

This program is a Fibonacci series generator written in Go. It calculates the Fibonacci series up to a user-specified limit. The key feature of this program is its use of the `math/big` package to handle large numbers, avoiding the integer overflow issues common in standard integer types.

## Features

- **Handle Large Numbers**: Utilizes `big.Int` from the `math/big` package to calculate large Fibonacci numbers efficiently.
- **User Input**: Prompts users to enter the number of Fibonacci series elements they want to generate.
- **Robust Input Validation**: Includes input validation to ensure the user enters a valid integer.

## Requirements

- Go (version 1.14 or later)

## Installation

1. Ensure Go is installed on your system. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).
2. Clone this repository or download the source code.

## Usage

1. Navigate to the directory containing the program.
2. Run the program using the Go command:

   ```bash
   go run main.go
   ```

3. Enter the number of elements you want in your Fibonacci series when prompted.

## How It Works

The program consists of two main functions:

- `fibonacci(limit int) []*big.Int`: This function generates a Fibonacci series up to the specified limit. It returns a slice of pointers to `big.Int` objects, each representing a number in the Fibonacci series.

- `getInput() int`: This function prompts the user to enter an integer value, which determines how many elements of the Fibonacci series to generate. It includes error handling to ensure valid input.

After the series is generated, it is printed to the console.

## Example Output

```plaintext
Enter an integer value: 10
Generating the series for : 10  numbers
The generated series for  10  numbers is  [0 1 1 2 3 5 8 13 21 34]
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
