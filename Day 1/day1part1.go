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
	mostCalories := findMostCalories(calorieList)
	fmt.Println(mostCalories)
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

func findMostCalories(list []int) int {
	max := 0
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}
