package main

import (
	"fmt"
	"strings"
)

func main() {
	// Task A
	unformattedPath := "/.././dffg/ddsd/ //sdfd/"

	fmt.Println(simplifyPath(unformattedPath))

	// Task B
	nums := []int{2, 2, 2, 2, 2}
	target := 8

	res := fourSum2(nums, target)

	fmt.Println("Fours is:", res)

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

func fourSum2(nums []int, target int) [][]int {

	var result [][]int
	pairs := findPairs(nums)
	triples := findTriples(nums, pairs)
	fours := findFours(nums, triples)

	for _, four := range fours {
		var sum int
		for _, index := range four {
			sum += nums[index]
		}

		if sum == target {
			resultFour := []int{nums[four[0]], nums[four[1]], nums[four[2]], nums[four[3]]}
			if !isDuplicate(result, resultFour) {
				result = append(result, resultFour)
			}
		}
	}

	return result
}

func findPairs(nums []int) (pairs [][]int) {

	for i, _ := range nums {
		for j, _ := range nums {
			var pair []int
			pair = append(pair, i)
			if j > i {
				pair = append(pair, j)
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}

func findTriples(nums []int, pairs [][]int) (triples [][]int) {
	for _, pair := range pairs {
		for j, _ := range nums {
			var triple []int
			triple = pair
			if j > pair[1] {
				triple = append(triple, j)
			}
			if len(triple) == 3 {
				triples = append(triples, triple)
			}
		}
	}
	return triples
}

func findFours(nums []int, triples [][]int) [][]int {
	var fours [][]int
	for _, triple := range triples {
		for j, _ := range nums {
			var four = make([]int, 3)
			copy(four, triple)
			if j > triple[2] {
				four = append(four, j)
				fours = append(fours, four)
			}

		}
	}
	return fours
}

func isDuplicate(result [][]int, second []int) bool {
	for _, four := range result {
		if four[0] == second[0] && four[1] == second[1] && four[2] == second[2] && four[3] == second[3] {
			return true
		}
	}
	return false
}
