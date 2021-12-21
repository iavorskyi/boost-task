package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {

	unformattedPath := "/.././dffg/ddsd/ //sdfd/"
	fmt.Println(simplifyPath(unformattedPath))

	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	res := fourSum(nums, target)

	fmt.Println("Result is:", res)

}

func simplifyPath(unformattedPath string) string {
	if unformattedPath[0] != '/' {
		unformattedPath = string('/') + unformattedPath
	}
	var fixedPath1 string

	for i, r := range unformattedPath {
		if r == ' ' {
			continue
		}
		if i != 0 && r == '/' && r == rune(unformattedPath[i-1]) {
			continue
		}
		fixedPath1 += string(r)

	}
	var result string

	splitedPath := strings.Split(fixedPath1, "/")
	for i, dir := range splitedPath {
		if dir == "" {
			continue
		}
		if dir == "." || dir == ".." {
			continue
		}
		if splitedPath[i+1] == ".." && i != 0 {
			continue
		}
		result += "/" + dir
	}
	if result == "" {
		result = "/"
	}

	return result
}

func fourSum(nums []int, target int) [][]int {
	if target > int(math.Pow10(9)) || target < -int(math.Pow10(9)) {
		fmt.Println("Illegal target:", target)
	}
	var result [][]int

	for {
		var usedIndexes []int
		var sum int
		if len(nums) < 4 {
			break
		}
		for len(usedIndexes) < 4 {

			usedIndexesUpdated, err := findNextNum(usedIndexes, nums)
			if err != nil {
				log.Println("wrong loop")
			}
			usedIndexes = usedIndexesUpdated
		}
		var resultFour []int
		for _, index := range usedIndexes {
			sum += nums[index]
			resultFour = append(resultFour, nums[index])
		}
		if sum == target {
			var isDoublicate bool
			log.Println("Sum:", sum)
			for _, four := range result {
				if four[0] == resultFour[0] && four[1] == resultFour[1] && four[2] == resultFour[2] && four[3] == resultFour[3] {
					isDoublicate = true
				}
			}
			if !isDoublicate {
				result = append(result, resultFour)
			}
		}
		nums = nums[1:]
		log.Println("Length of nums:", len(nums))
	}

	return result
}

func findNextNum(usedIndexes []int, nums []int) (usedIndexesUpdated []int, err error) {
	log.Println("In findNextNum", usedIndexes)
	usedIndexesUpdated = usedIndexes
	for i, _ := range nums {
		if !contains(usedIndexes, i) {
			usedIndexesUpdated = append(usedIndexes, i)
			return usedIndexesUpdated, nil
		}
	}
	return usedIndexesUpdated, errors.New("no suetable nums")
}

func contains(s []int, i int) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
