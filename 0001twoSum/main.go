
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		_, ok := m[another]
		if ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
func main() {
	var b = []int{1,2,3,34,566,2,23,455,2,23}
	var a = 3
	fmt.Println(twoSum(b, a))
}
