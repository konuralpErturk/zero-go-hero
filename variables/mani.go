package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//lessonVariables()

	//lessonConsoleInput()

	lessonMathOperations()
}

func lessonMathOperations() {
	i1, i2, i3 := 12, 24, 36
	f1, f2, f3 := 12.2, 24.4, 36.6

	sumInt := i1 + i2 + i3
	sumFloat := f1 + f2 + f3

	println("Integers Sum: ", sumInt)

	println("Floats Sum: ", sumFloat)

	total := float64(i1) + f1 + float64(i2)

	println("total: ", total)

}

func lessonConsoleInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter a Number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of Number", f)
	}
}

func lessonVariables() {
	str1 := "konuralp01"
	str2 := "konuralp02"
	str3 := "konuralp03"

	aNumber := 42

	fmt.Println(str1, str2, str3)
	fmt.Println(aNumber)

	stringLenght, err := fmt.Println("This value is", aNumber)
	if err == nil {
		fmt.Println("String length:", stringLenght)
	}

	fmt.Printf("Value of number: %v\n", aNumber)

	fmt.Printf("Data type: %T\n", aNumber)
}
