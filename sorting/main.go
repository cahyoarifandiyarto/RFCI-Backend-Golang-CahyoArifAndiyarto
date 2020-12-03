package main

import "fmt"

func sort(arr []int) {
	total := 0
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				total++
				fmt.Println("[", arr[j-1], ",", arr[j], "]", "->", arr)
			}
			j = j - 1
		}
	}
	fmt.Println("Jumlah swap:", total)
}

func main() {
	arr := []int{4, 9, 7, 5, 8, 9, 3}
	sort(arr)
}
