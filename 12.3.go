package main

import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func findMedian(arr []int) int {
	n := len(arr)

	if n%2 != 0 {
		return arr[n/2]
	}

	return (arr[(n-1)/2] + arr[n/2]) / 2
}

func main() {
	var num int
	var data []int

	for {
		fmt.Scan(&num)

		if num == -5313 {
			break
		}

		if num == 0 {
			if len(data) > 0 {
				sortedData := insertSort(append([]int{}, data...))
				median := findMedian(sortedData)
				fmt.Println(median)
			}
		} else {

			data = append(data, num)
		}
	}
}
