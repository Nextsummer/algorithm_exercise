package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSumHash(nums, 6))
}

func twoSumHash(nums []int, target int) []int {
	numMap := make(map[int]int, 2)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{i, j}
		}
		numMap[num] = i
	}
	return nil
}

func twoSumDeRand(nums []int, target int) []int {
	rand.Seed(time.Now().Unix())
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumRand(nums []int, target int) []int {
	rand.Seed(time.Now().Unix())
	length := len(nums)
	for i := 0; i < 10*length*length; i++ {
		p1 := rand.Intn(length)
		p2 := 0

		for {
			p2 = rand.Intn(length)
			if p1 == p2 {
				continue
			}
			break
		}
		if nums[p1]+nums[p2] == target {
			return []int{p1, p2}
		}
	}
	return nil
}
