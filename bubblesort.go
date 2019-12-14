package main

/* Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. 
The program should print the integers out on one line, in sorted order, from least to greatest. 
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument
 and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position 
of two adjacent elements in the slice. You should write a Swap() function which performs this operation. 
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a 
position in the slice. The Swap() function should return nothing, but it should swap the contents
 of the slice in position i with the contents in position i+1.
   */
   
import (
	"fmt"
	"strings"
	"strconv"
)


func main() {
	var strFromUser string

	fmt.Printf("enter sequence of up to 10 integers (seperated by spaces): ")
	n, err := fmt.Scanf("%q", &strFromUser)
	fmt.Printf("Got %d characters\n", n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strFromUser)
	sliceofInts := make([]int,0,10)
	sliceOfStrs := strings.Split(strFromUser, " ")
	fmt.Println(sliceOfStrs)
	for _, s := range sliceOfStrs {
		newInt, _ := strconv.Atoi(s)
		sliceofInts = append(sliceofInts, newInt)
	}
	fmt.Println(sliceofInts)
	Swap(sliceofInts, 3)
	fmt.Println(sliceofInts)
	fmt.Println("now sort the whole thing")
	BubbleSort(sliceofInts)
}

func BubbleSort(sliceToSort []int){
	for i, _:= range sliceToSort {
		if !sweep(sliceToSort, i) {
			fmt.Println(sliceToSort)
			return
		}
	}
}

func sweep(sliceToSort []int, prevPasses int) bool {
	didSwap := false
	for i := 0; i < len(sliceToSort) - prevPasses; i++ {
		if len(sliceToSort) > i+1 {
			fmt.Println("comparing: ", sliceToSort[i], sliceToSort[i+1])
			if sliceToSort[i] > sliceToSort[i+1] {
				Swap(sliceToSort, i)
				didSwap = true
			}
		}
	}
	return didSwap
}

func Swap(sints []int, indx int) {
	// swap the int at position indx with the one at indx+1
	intToSwap := sints[indx]
	sints[indx] = sints[indx+1]
	sints[indx+1] = intToSwap
	fmt.Println("Swapped: ", sints[indx], intToSwap)
}
