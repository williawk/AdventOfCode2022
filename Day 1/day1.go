package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
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

	tempStorage := []int{}
	calories := []int{}

	for scanner.Scan() { // internally, it advances token based on separator
		if scanner.Text() == "" {
			var sumCal int
			for i := 0; i >= len(tempStorage)-1; i++ {
				sumCal = sumCal + tempStorage[i]
			}
			calories = append(calories, sumCal)
			tempStorage = []int{}
		}
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		tempStorage = append(tempStorage, number)
	}
}
