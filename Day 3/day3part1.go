package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rucksackInventories := importRucksacks()
	commonItems := examineRucksacks(rucksackInventories)
	priorityMap := buildPriorityMap()
	prioritySum := calcPrioritySum(commonItems, priorityMap)
	fmt.Println(prioritySum)
}

func importRucksacks() []string {
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

func examineRucksacks(list []string) []string {
	commonItems := []string{}
	for i := 0; i < len(list); i++ {
		splitIndex := (len(list[i]) / 2)
		var compartment1 string = string(list[i][0:splitIndex])
		var compartment2 string = string(list[i][splitIndex:len(list[i])])

		commons := checkForCommons(compartment1, compartment2)
		commonItems = append(commonItems, commons)
	}
	return commonItems
}

func checkForCommons(comp1 string, comp2 string) string {
	var commons string
	for i := 0; i < len(comp1); i++ {
		if strings.Contains(comp2, string(comp1[i])) {
			commons = string(comp1[i])
			break
		}
	}
	return commons
}

func buildPriorityMap() map[string]int {
	const abcABC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorityMap := make(map[string]int)

	for i := 1; i < 53; i++ {
		priorityMap[string(abcABC[i-1])] = i
	}
	return priorityMap
}

func calcPrioritySum(commonItems []string, priorityMap map[string]int) int {
	prioritySum := 0

	for i := 0; i < len(commonItems); i++ {
		prioritySum += priorityMap[commonItems[i]]
	}
	return prioritySum
}
