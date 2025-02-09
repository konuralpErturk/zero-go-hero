package main

import (
	"fmt"
)

func main() {
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
}
