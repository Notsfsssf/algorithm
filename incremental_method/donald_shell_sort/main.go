package main

import "fmt"

//希尔排序
func main() {
	arr := []int{13, 14, 94, 33, 82, 25}
	shellSort(arr, len(arr)/2)
	fmt.Println(arr)
}

/**
希尔排序：把切片分成n个batch，对每个batch进行插入排序；然后减小batch，再对每个batch进行插入排序；直到bathc等于1
*/
func shellSort(arr []int, gap int) {
	if gap < 1 {
		return
	}
	// k : 每个batch中的元素所在batch的index， 介于[0, batchSize)
	for k := 0; k < gap; k++ {
		// 用到了插入排序
		for j := 1; gap*j+k < len(arr); j++ { // j: 用来获取每个batch所在的第k个元素，拥有多少个batch
			for n := j; n > 0; n-- {
				pre := gap*(n-1) + k
				next := gap*n + k
				if arr[next] < arr[pre] {
					arr[next], arr[pre] = arr[pre], arr[next]
				}
			}

		}
	}
	shellSort(arr, gap/2)
}
