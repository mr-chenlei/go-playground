package main

import "log"

type someStruct struct {
	A string
	B string
	C string
}

func returnStruct() *someStruct {
	s := &someStruct{
		A: "A",
	}
	return s
}

func main() {
	ss := &someStruct{
		B: "B1",
		C: "C1",
	}
	log.Println(ss)

	ss = returnStruct()
	ss.B = "B2"
	ss.C = "C2"
	log.Println(ss)
}
