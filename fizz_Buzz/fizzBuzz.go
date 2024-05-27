package module01

import (
	"strconv"
)

func fizzBuzz(number int) []string {
	newNum := number + 1
	arr := make([]string, newNum)
	for i := 1; i < newNum; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr[i] = "Fizz Buzz,"
		} else if i%3 == 0 {
			arr[i] = "Fizz,"
		} else if i%5 == 0 {
			arr[i] = "Buzz,"
		} else {
			arr[i] = strconv.Itoa(i) + ","
		}
	}
	return arr
}
