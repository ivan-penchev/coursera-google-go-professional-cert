package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, v0, s0, t float64

	a = ReadVariable("Enter acceleration:", "value for acceleration is incorrect please enter a valid number:")
	v0 = ReadVariable("Enter  initial velocity:", "value for initial velocity is incorrect please enter a valid number:")
	s0 = ReadVariable("Enter  initial displacement:", "value for  initial displacement is incorrect please enter a valid number:")
	t = ReadVariable("Enter a value for time:", "value for time is incorrect please enter a valid number:")
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Displacement travelled: ", fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + v0*t + s0
	}
}

func ReadVariable(message string, errorMessage string) float64 {
	var error error
	var number float64
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(message)
		input, _ := reader.ReadString('\n')
		acceleration := input[0 : len(input)-2]
		number, error = strconv.ParseFloat(acceleration, 64)
		if error == nil {
			break
		}
		fmt.Println(errorMessage)
	}
	return number
}
