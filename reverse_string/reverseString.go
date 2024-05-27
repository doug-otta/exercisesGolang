package reversestring

func StringReverse(word string) string {
	var reversestring string
	for i := len(word) - 1; i >= 0; i-- {
		reversestring += string(word[i])
	}
	return reversestring

}
