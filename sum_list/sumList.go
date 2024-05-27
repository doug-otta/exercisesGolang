package sumlist

func sumList(list []int) int {
	iterate := 0
	for _, value := range list {
		iterate += value
	}
	return iterate
}
