package main

import "fmt"

func main() {
	fmt.Println("No. of digits")
	var num int
	fmt.Scanln(&num)

	arr := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("enter digits")
		var digits int
		fmt.Scanln(&digits)
		arr[i] = digits
	}
	fmt.Println("input: ", arr)

	var output = []int{}
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum
		output = append(output, sum)

	}
	fmt.Println("output : ", output)
}
