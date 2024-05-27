package main

import "fmt"

func findNumber(list []int, num int) bool {
	var findedNUm bool
	for _, value := range list {
		if value == num {
			findedNUm = true
			break
		}
	}
	return findedNUm
}

func main() {
	check := findNumber([]int{1, 2, 3, 4, 5, 6}, 3)
	fmt.Println(check)
}
