package iteration

import "fmt"

func Iteration(num int, letter string) string {
	concat := ""
	for i := 0; i < num; i++ {
		concat += letter
	}

	fmt.Println(concat)

	return concat
}
