package main
import "fmt"
import "leetcode/Sort/utils"

func main() {

    list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 9233, 2222, 4567}
    utils.InsertSort(list2)
    fmt.Println("直接插入排序", list2)
}
