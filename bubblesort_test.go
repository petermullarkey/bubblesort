package main

/* Test the program by running it twice and testing it with a different sequence of integers each time. 
The first test sequence of integers should be all positive numbers 
and the second test should have at least one negative number.
*/

import ("testing"
		"reflect"
)

func Test(t *testing.T) {
     var tests = []struct {
     	 s []int; want []int
	}{
		{[]int{9, 2, 6, 7, 3, 8}, []int{2, 3, 6, 7, 8, 9}},
		{[]int{9, 2, 6, 1, -2, 8}, []int{-2, 1, 2, 6, 8,9}},
	}
	for _, c := range tests {
		BubbleSort(c.s)
	    if !reflect.DeepEqual(c.s, c.want) {
	       t.Errorf("BubbleSort(%q) wanted %q", c.s, c.want)
	    }
	}
}