package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greetings(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HandlerMyGreetings(w http.ResponseWriter, r *http.Request) {
	Greetings(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMyGreetings))

	if err != nil {
		fmt.Println(err)
	}
}
