package main

import "fmt"

func main() {
    // Integer data type
    var age int = 30
    fmt.Println("Age:", age)

    // String data type
    var name string = "John"
    fmt.Println("Name:", name)

    // Float data type
    var weight float64 = 70.5
    fmt.Println("Weight:", weight)

    // Boolean data type
    var isStudent bool = true
    fmt.Println("Is student?", isStudent)

    // Operations with integers
    num1 := 10
    num2 := 20
    fmt.Println("Sum:", num1+num2)
    fmt.Println("Difference:", num2-num1)
    fmt.Println("Product:", num1*num2)
    fmt.Println("Division:", num2/num1)

    // Operations with floats
    num3 := 3.14
    num4 := 1.5
    fmt.Println("Sum:", num3+num4)
    fmt.Println("Difference:", num3-num4)
    fmt.Println("Product:", num3*num4)
    fmt.Println("Division:", num3/num4)

    // String concatenation
    greeting := "Hello"
    subject := "World"
    fmt.Println(greeting + ", " + subject)
}
