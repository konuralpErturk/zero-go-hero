package main

import (
	"fmt"
	"sort"
)

func main() {
	//lessonPointer()

	//lessonArrays()

	//lessonSlice()

	//lessonMap()

	//lessonObject()
}

func lessonObject() {
	kangal := Dog{"Kangal", 31}
	fmt.Println(kangal)
	fmt.Printf("Breed name:%v,Weight:%v\n", kangal.Breed, kangal.Weight)
	kangal.Weight = 20
	fmt.Printf("Breed name:%v,Weight:%v\n", kangal.Breed, kangal.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}

func lessonMap() {
	states := make(map[string]string)
	states["41"] = "kocaeli"
	states["01"] = "adana"
	states["16"] = "bursa"

	fmt.Println(states)

	kocaeli := states["41"]

	fmt.Println(kocaeli)

	//delete
	delete(states, "16")
	fmt.Println(states)
	//add new
	states["06"] = "ankara"
	//for loop
	for k, v := range states {
		fmt.Printf("%v,%v \n", k, v)
	}
	//assigne to array
	stateList := make([]string, len(states))
	i := 0
	for k := range states {
		stateList[i] = k
		i++
	}
	fmt.Println(stateList)
	//sorted
	sort.Strings(stateList)
	//write console map sorted map with
	for i := range stateList {
		fmt.Println(states[stateList[i]])
		i++
	}

}

func lessonSlice() {
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Blue", "Yellow")
	fmt.Println(colors)
	fmt.Println(colors[1])
	colors = append(colors, "Orange", "Gray")
	fmt.Println(colors)
	fmt.Println(colors[4])

	sort.Strings(colors)
	fmt.Println(colors)

	colors = remove(colors, 0)

	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func lessonArrays() {
	var colors [3]string
	var numbers = [5]int{1, 2, 3, 4, 5}

	colors[0] = "Red"
	colors[1] = "Yellow"
	colors[2] = "black"

	fmt.Println("All colors is: ", colors)
	fmt.Println("Red color is: ", colors[0])
	fmt.Println("All numbers is: ", numbers)
	fmt.Println("1 numbers is: ", numbers[0])
}

func lessonPointer() {
	varInt := 42
	var p *int

	p = &varInt

	if p == nil {
		fmt.Println("p value is: ", p)
	} else {
		fmt.Println("p value is: ", *p)
		fmt.Println("p reference is: ", p)
	}

	value1 := 43
	pointer1 := &value1
	*pointer1 = *pointer1 / 10
	fmt.Println("pointer1 value is:", *pointer1)
	fmt.Println("pointer1 pointer is:", pointer1)
}
