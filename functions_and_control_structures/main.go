package main

import (
	"errors"
	"fmt"
)

func main() {
	value := "Zira"
	printString(value)

	fmt.Printf("Returned Function \n The result is %v", intMult(2, 3))
	var a int = 10
	var b int = 5
	// var b_error int = 0
	result1, result2 := intDiv(a, b)

	fmt.Printf("\nReturning Multiple Values \n The result is  of Division is %v wih remainder %v", result1, result2)

	var result, remainder, err = intDivErr(a, b)
	if err != nil { //default value of error is nil
		fmt.Println("\nError:", err)
	} else if remainder == 0 {
		fmt.Printf("\nError Handling \n The result of division is %v", result)
	} else {
		fmt.Printf("\nError Handling \n The result of division is %v with remainder %v", result, remainder)
	}

	fmt.Println("\nSwitch Conditional")
	switch remainder {
	case 0:
		fmt.Println("The remainder was exact")
	case 1, 2:
		fmt.Println("The remainder was close")
	default:
		fmt.Println("The remainder was not close")
	}
}
func printString(value string) {
	fmt.Printf("Hello,%v!", value)
}

// Returned Function
func intMult(a int, b int) int {
	var result = a * b
	return result
}

// Returning Multiple Values
func intDiv(a int, b int) (int, int) {
	var result1 = a / b
	var result2 = a % b
	return result1, result2
}

// Error Handling
func intDivErr(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("cannot do division by zero")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
