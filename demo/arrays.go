package main

import "fmt"

func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,5,6,7,8}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
