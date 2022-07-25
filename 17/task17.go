package main

import "fmt"

func main() {
	mass := []int{10, 18, 54, 65, -11, 12, 36, 185, 8, 3, 92, -17, -6}
	quickSort(mass, 0, len(mass)-1)
	//Массив должен быть упорядоченным
	fmt.Println(mass)
	target := 12
	resultIndex := binarySearch(mass, target)
	if resultIndex >= 0 {
		fmt.Printf("Value: %d, index: %d", target, resultIndex)
	} else {
		fmt.Printf("Value %d not found", target)
	}
}

func binarySearch(data []int, target int) int {
	left := 0
	right := len(data) - 1
	var middle int

	for left <= right {
		middle = left + (right-left)/2
		if target == data[middle] {
			return middle
		}
		if target < data[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
