package main

import "fmt"

//堆排序
func main() {
	array := []int{492, 38, 65, 97, 487, 11, 55, 38, 54848, 88, 447}
	heapSort(array)
	fmt.Println(array)
}

func heapSort(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjustHeap(array, i, len(array))
	}
	for j := len(array) - 1; j > 0; j-- {
		swap(array, 0, j)
		adjustHeap(array, 0, j)
	}
}

func adjustHeap(array []int, i int, length int) {
	temp := array[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 { //原来如此，我现在完全搞懂了.jpg
		if k+1 < length && array[k] < array[k+1] {
			k++
		}
		if array[k] > temp {
			swap(array, i, k)
			i = k
		} else {
			break
		}
	}

}
func swap(array []int, a, b int) {
	array[a], array[b] = array[b], array[a]
}
