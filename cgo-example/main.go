package main

/*
int add2(int a, int b) {
	return a + b;
}
*/
import "C"
import (
	_ "unsafe"
)

func go_add2(a, c int) int {
	return a + c
}

func GoAdd(num int) {
	for i := 0; i < num; i++ {
		go_add2(i, i)
	}
}

func CAdd(num int) {
	for i := 0; i < num; i++ {
		C.add2(C.int(i), C.int(i))
	}
}

func main() {

}
