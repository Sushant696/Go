package main

import "fmt"

// Function to calculate the number of digits in a given number
func countDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

// Function to check if a number is an Armstrong number
func isArmstrong(number int) bool {
	// Calculate the number of digits in the given number
	numDigits := countDigits(number)

	// Temporary variable to store the original number
	temp := number

	// Variable to store the sum of the powers of individual digits
	sum := 0

	// Calculate the sum of powers of individual digits
	for temp != 0 {
		digit := temp % 10
		sum += power(digit, numDigits)
		temp /= 10
	}

	// Check if the sum equals the original number
	return sum == number
}

// Function to calculate power of a number without using math.Pow
func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// creating an function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	num1 := 153
	fmt.Printf("Is %d an Armstrong number? %t\n", num1, isArmstrong(num1))

	// while making array we have to first give the number of items that we need that we need and then we have to specify the data type

	// Array with integer
	// arr2 := [5]int{5, 6, 7}

	// Array with string
	arr := []string{"hello", "world", "trying go lang."}
	// fmt.Println(arr2)
	// fmt.Println(arr)
	arr = append(arr, "new element")
	fmt.Println("Slice after appending:", arr)
	// use loop to iterate over an array
	length := len(arr)
	for i := 0; i < length; i++ {
		arr = append(arr, arr[i])
		println(arr[i])
	}

	fmt.Println("Slice after appending:", arr)

	var num int
	fmt.Print("Enter a number to find factorial: ")
	fmt.Scanln(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))

	fmt.Println("Using for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Using for loop as a while loop to print numbers from 1 to 5
	fmt.Println("Using for loop as while loop:")
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	for i := 1; i <= 100; i++ {
		// Check if current number is divisible by both 3 and 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
