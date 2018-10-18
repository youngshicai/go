package main

import "fmt"

func partition(arr []int, len int) int { //len >=2
	minIndex := 0
	i := 0
	pivot := arr[len-1]

	for i < len-1 {
		if arr[i] < pivot {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
			minIndex++
		}
		i++
	}

	arr[minIndex], arr[i] = arr[i], arr[minIndex]

	//fmt.Println(arr)
	return minIndex

}

func quickSort(arr []int) {

	var q int
	length := len(arr)
	if length >=2 {
		q = partition(arr, length)
	}else{
		return
	}

	if q >= 2 {//长度比较，这里完全可以不要的，q>=0,arr[:q]不会报错
		quickSort(arr[:q])//半开半闭区间别忘了
	}

	if length-q-1 >= 2 {//长度比较，这里必须要的，q+1溢出切片长度（q+1 =len）,防止编译错误
		quickSort(arr[q+1:])
	}

	return

}

func main() {
	a := []int{1, 35, 546, 76, 2, 9, 7, 5, 4, 3, 2, 1, 8, 6, 23, 546, 786, 880, 4355, 100, 23, 34, 45, 32, 12, 23, 12, 11, 123, 453, 67, 787, 990}

	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}
