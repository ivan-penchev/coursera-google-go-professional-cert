package main

import (
	"fmt"
	s "strings"
)

func main() {
	var user_input string
	fmt.Println("Give me string:")
	fmt.Scan(&user_input)
	if s.HasPrefix(user_input, "i") && s.Contains(user_input, "a") && s.HasSuffix(user_input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
