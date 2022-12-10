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
	for i := 0; i < len(list); i += 3 {
		commons := checkForCommons(list[i], list[i+1], list[i+2])
		commonItems = append(commonItems, commons)
	}
	return commonItems
}

func checkForCommons(bag1 string, bag2 string, bag3 string) string {
	var commons []string
	var common string
	for i := 0; i < len(bag1); i++ {
		if strings.Contains(bag2, string(bag1[i])) {
			commons = append(commons, string(bag1[i]))
		}
	}
	for i := 0; i < len(commons); i++ {
		if strings.Contains(bag3, string(commons[i])) {
			common = string(commons[i])
			break
		}
	}
	return common
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
