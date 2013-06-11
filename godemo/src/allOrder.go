package main

import (
	"fmt"
)

//数组全排列算法
//使用递归方法
//

var aXXX = []int{2, 4, 5, 8,9,7}

func sortME(sourcArr []int, ipos int, outArr *[][]int) {
	//fmt.Println(sourcArr)

	if len(sourcArr) <= 2 {
		var tmpArr = [][]int{{sourcArr[0], sourcArr[1]}, {sourcArr[1], sourcArr[0]}}

		*outArr = append(*outArr, tmpArr[0], tmpArr[1])

		return

	}

	for i := 0; i < len(sourcArr); i++ {
		leader := sourcArr[i]

		tmpOutArr := make([][]int, 0)
		sortME(delElement(sourcArr, i), i, &tmpOutArr)
		//      fmt.Println(outArr)

		for z := 0; z < len(tmpOutArr); z++ {
			
			*outArr = append(*outArr, leadElement(tmpOutArr[z], leader)) 
		}

	}

	return

}

func leadElement(sArr []int, val int) []int {
	retArr := make([]int, len(sArr)+1)
	retArr[0] = val

	j := 1
	for i := 0; i < len(sArr); i++ {
		retArr[j] = sArr[i]
		j++

	}
	return retArr

}

func delElement(sArr []int, iPos int) []int {
	retArr := make([]int, len(sArr)-1)

	j := 0
	for i := 0; i < len(sArr); i++ {
		if i != iPos {
			retArr[j] = sArr[i]
			j++
		}

	}
	return retArr

}

func main() {

	outArray := make([][]int, 0)

	sortME(aXXX, 0, &outArray)

	fmt.Println(len(outArray))
	fmt.Println(outArray)

}
