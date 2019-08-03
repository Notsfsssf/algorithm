package main

import "errors"

//查找
//参考http://codingxiaxw.cn/2017/01/14/66-leetcode-find/
type Search struct {
	key   int
	array []int
}
type Tree struct {
	value int
	left  Tree
	right Tree
}

func NewSearch(key int, array []int) *Search {
	return &Search{key: key, array: array}
}

func main() {

}

//顺序
func (s *Search) orderSearch() (int, error) {
	if s.array == nil || len(s.array) < 1 {
		return -1, errors.New("-1")
	}
	for i := 0; i < len(s.array); i++ {
		if s.array[i] == s.key {
			return i, nil
		}

	}
	return -1, errors.New("Not found")

}

//二分
func (s *Search) binarySearch() (int, error) {
	start, end := 0, len(s.array)-1
	for start <= end {
		mid := (start + end) / 2
		if s.key == s.array[mid] {
			return mid, nil
		} else if s.key < s.array[mid] { //大小限定
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1, errors.New("not found")
}
