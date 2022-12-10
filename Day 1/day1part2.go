package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputList := createInputList()
	calorieList := createCalorieList(inputList)
	top3Calories := findTop3Calories(calorieList)
	sumTop3 := sumTop3Values(top3Calories)

	fmt.Println(sumTop3)
}

func createInputList() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	list := []string{}

	for scanner.Scan() { // internally, it advances token based on separator
		list = append(list, scanner.Text())
	}

	return list
}

func createCalorieList(inputList []string) []int {
	tempStorage := []int{}
	calories := []int{}

	for i := 0; i < len(inputList); i++ {
		if inputList[i] == "" {
			sum := sumCalories(tempStorage)
			calories = append(calories, sum)
			tempStorage = []int{}
			continue
		}
		number := makeNumber(inputList[i])
		tempStorage = append(tempStorage, number)
	}

	return calories
}

func sumCalories(list []int) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func makeNumber(x string) int {
	number, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func findTop3Calories(list []int) [3]int {
	top3 := [3]int{list[0], list[1], list[2]}

	for i := 3; i < len(list); i++ {
		lowestNum, index := findLowestTop3(top3)

		if list[i] > lowestNum {
			top3[index] = list[i]
		}
	}
	return top3
}

func findLowestTop3(list [3]int) (int, int) {
	lowest := list[0]
	lowIndex := 0
	if list[1] < lowest {
		lowest = list[1]
		lowIndex = 1
	}
	if list[2] < lowest {
		lowest = list[2]
		lowIndex = 2
	}
	return lowest, lowIndex
}

func sumTop3Values(list [3]int) int {
	sum := list[0] + list[1] + list[2]
	return sum
}
