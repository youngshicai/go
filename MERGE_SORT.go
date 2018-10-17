package main

import "fmt"

func heBin(arrA []int, arrB []int) ([]int) {
	i := 0
	j := 0
	lenA := len(arrA)
	lenB := len(arrB)
	res := make([]int,0, lenA+lenB)

	for i < lenA || j < lenB {

		if i == lenA {
			res = append(res, arrB[j:]...)
			break
		} else if j == lenB {
			res = append(res, arrA[i:]...)
			break
		}

		if arrA[i] < arrB[j] {
			res = append(res, arrA[i])
			i++
		} else {
			res = append(res, arrB[j])
			j++
		}
	}

	return res


}

func guiBin(arr []int) ([]int) {
	q := len(arr) / 2

	if q < 1 {
		return arr[:1]
	}

	arrA := guiBin(arr[:q])
	arrB := guiBin(arr[q:])

	res := heBin(arrA, arrB)

	return res

}

func main() {
	a := []int{1, 35, 546, 76,2,9,7,5,4,3,2,1, 8, 6, 23, 546, 786, 880, 4355, 100,23,34,45,32,12,23,12,11,123, 453, 67, 787, 990}
	b := guiBin(a)

	fmt.Println(a)
	fmt.Println(b)
}
