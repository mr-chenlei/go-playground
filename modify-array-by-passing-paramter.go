package main

import "fmt"

type Data struct {
	i int
}

func modifyArray(arr []*Data) []*Data{
	for i := 10; i < 20; i++ {
		d := &Data{i}
		arr = append(arr, d)
	}
	return arr
}

func main() {
	arr := make([]*Data, 0)
	for i := 0; i < 10; i++ {
		d := &Data{i}
		arr = append(arr, d)
	}
	// 1.Before modification
	fmt.Println("Before modification...")
	for _, v := range arr {
		fmt.Printf("%d ", v.i)
	}
	fmt.Println()
	// 2.After modification
	anotherArr := modifyArray(arr)
	fmt.Println("After modification...")
	for _, v := range arr {
		fmt.Printf("%d ", v.i)
	}
	fmt.Println()
	// 3.Another copy of arr
	fmt.Println("Another copy of arr modification...")
	for _, v := range anotherArr {
		fmt.Printf("%d ", v.i)
	}
	fmt.Println()
}