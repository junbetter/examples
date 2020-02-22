package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(4, 7))
}

func add(a, b int) int {
	for b != 0 {
		tmp := a ^ b     //相加各位的值, 不计进位. 相同位置0, 相反位置1
		b = (a & b) << 1 //计算进位的值, 先保留同为1的位, 都为1的位要向左进位, 因此左移1位
		a = tmp          //让a作为tmp(相加各位的值的结果), 再和b(进位)去异或, 就是相加操作, 只要进位不为0, 这个循环就一直进行下去
	}
	return a
}
