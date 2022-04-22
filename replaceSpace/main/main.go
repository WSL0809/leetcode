package main

import "fmt"
import "leetcode/replaceSpace/utils" 
//这一题替换空格，一个方法是建立一个新的字符串，用for range遍历原字符串，使用Switch对v值进行选择，case空格则加上%20，
//默认情况则是直接加上v，注意要转为string即可。



func main() {
	
	fmt.Println(utils.ReplaceSpace02("wo 们 是 一 家 人"))
}