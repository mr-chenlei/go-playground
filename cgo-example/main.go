package main

import (
	_ "unsafe"
)

/*
int add2(int numb, int a, int b) {
	int sum;
	for(int i = 0; i < numb; i++)
	{
		sum += (a + b);
	}
	return sum;
}
*/
import "C"

func goAdd2(a, c int) int {
	return a + c
}

func GoAdd(num int) {
	for i := 0; i < num; i++ {
		goAdd2(i, i)
	}
}

func CAdd(num int) {
	C.add2(C.int(num), C.int(1), C.int(1))
}

func main() {
}
