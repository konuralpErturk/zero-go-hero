package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	value1 := 5
	value2 := 10
	value3 := 56

	sum := getSum(value1, value2)
	fmt.Printf("The sum of %v and %v is %v\n", value1, value2, sum)

	sum2 := gelAllSum(value1, value2, value3)
	fmt.Printf("The sum of %v, %v and %v is %v\n", value1, value2, value3, sum2)

	sum, count, avarage := getAllValues(value1, value2, value3)
	fmt.Printf("The sum of %v, %v and %v is %v\n", value1, value2, value3, sum)
	fmt.Printf("The count of %v, %v and %v is %v\n", value1, value2, value3, count)
	fmt.Printf("The avarage of %v, %v and %v is %v\n", value1, value2, value3, avarage)

}

func getSum(value1, value2 int) int {
	return value1 + value2
}

func gelAllSum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func getAllValues(values ...int) (int, int, int) {
	sum := 0
	count := 0
	for _, value := range values {
		sum += value
		count++
	}
	avarage := sum / count
	return sum, count, avarage
}
