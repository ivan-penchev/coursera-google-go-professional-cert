package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input numbers(separate with space):")
	br := bufio.NewReader(os.Stdin)
	a, _, _ := br.ReadLine()
	ns := strings.Split(string(a), " ")
	var values []int
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	BubbleSort(&values)
	fmt.Println(values)
}

func BubbleSort(valuesToSort *[]int) {
	// Look at the first number in the list.
	// Compare the current number with the next number.
	// Is the next number smaller than the current number? If so, swap the two numbers around. If not, do not swap.
	// Move to the next number along in the list and make this the current number.
	// Repeat from step 2 until the last number in the list has been reached.
	// If any numbers were swapped, repeat again from step 1.
	// If the end of the list is reached without any swaps being made, then the list is ordered and the algorithm can stop.
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(*valuesToSort)-1; i++ {
			// if the current element is greater than the next element, swap them
			if (*valuesToSort)[i] > (*valuesToSort)[i+1] {
				// log that we are swapping values for posterity
				fmt.Println("Swapping")
				Swap(valuesToSort, i)
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}

}

func Swap(values *[]int, indexToSwap int) {
	//use Go's tuple assignment to swap values
	(*values)[indexToSwap], (*values)[indexToSwap+1] = (*values)[indexToSwap+1], (*values)[indexToSwap]
}
