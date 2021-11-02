package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type SortedPart struct {
	goroutineId int
	arr         []int
}

func Print(id int, arr []int) {
	fmt.Printf("goroutine #%d:", id)
	for _, n := range arr {
		fmt.Printf(" %d", n)
	}
	fmt.Println()
}

func Printer(wg *sync.WaitGroup, nParts int, ch chan SortedPart) {
	defer wg.Done()
	for i := 0; i < nParts; i++ {
		sp := <-ch
		Print(sp.goroutineId, sp.arr)
	}
}

func SortPart(id int, wg *sync.WaitGroup, arr []int, ch chan SortedPart) {
	defer wg.Done()
	sort.Ints(arr)
	sp := SortedPart{id, arr}
	ch <- sp
}

func MergeParts(a []int, b []int) []int {
	res := []int{}
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}

func Split(arr []int, nParts int) [][]int {
	res := [][]int{}
	partSz := len(arr) / nParts
	rem := len(arr) % partSz
	beg := 0
	for i := 0; i < nParts; i++ {
		var add int
		if rem > 0 {
			add = 1
			rem--
		}
		res = append(res, arr[beg:beg+partSz+add])
		beg += partSz + add
	}
	return res
}

func main() {
	nParts := 4
	arr := []int{}
	fmt.Println("Enter several numbers:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	if len(tokens) < nParts {
		fmt.Printf("Error: number of numbers should not be less than %d\n", nParts)
		return
	}
	for _, t := range tokens {
		n, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		arr = append(arr, n)
	}
	ch := make(chan SortedPart, nParts)
	var wg sync.WaitGroup
	wg.Add(5)
	go Printer(&wg, nParts, ch)

	parts := Split(arr, nParts)
	for i, p := range parts {
		go SortPart(i+1, &wg, p, ch)
	}

	wg.Wait()

	res := []int{}
	for _, p := range parts {
		res = MergeParts(res, p)
	}
	for _, n := range res {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
