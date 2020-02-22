package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(8))
}

//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//n个骰子的点数和的最小的值为n，最大值为6n, 所以点数和的情况会有6n-n+1种
//举例当n为3时，最小为3，最大为18，一共有16种和(包含3)
func solution(n int) []float64 {
	if n < 1 {
		return nil
	}
	max := 6
	t := n
	m := map[int]int{}

	//当n为1时每个点数出现的次数情况都为1
	for i := 1; i <= max; i++ {
		m[i] = 1
	}

	//当n大于1时需计算n个相加的点数出现的次数
	for t > 1 {
		tmp := map[int]int{}
		for i := 1; i <= max; i++ {
			for k, v := range m {
				tmp[k+i] += v
			}
		}
		m = tmp
		t--
	}

	total := math.Pow(float64(max), float64(n))
	ret := make([]float64, n*max-n+1)
	for k, v := range m {
		ret[k-n] = float64(v) / total
	}

	return ret
}
