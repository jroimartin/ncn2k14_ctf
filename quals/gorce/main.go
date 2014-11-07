package main

import (
	"fmt"
	"os"
)

func main() {
	a := []byte{0x65, 0x87, 0x33, 0x45, 0x12}
	// G0w1n!C0ngr4t5!!
	b := []byte{85, 117, 68, 182, 11, 51, 6, 3, 233, 2, 96, 113, 71, 178, 68, 51}

	c := []byte{0x92, 0x56, 0x10, 0x22, 0x81}
	// Yeah!
	d := []byte{203, 51, 113, 74, 160, 152}
	// Nope!
	e := []byte{220, 57, 96, 71, 160, 152}

	if len(os.Args) != 2 {
		for i := 0; i < len(e); i++ {
			fmt.Print(string(c[i%len(c)] ^ e[i]))
		}
		return
	}
	if len(os.Args[1]) != len(b) {
		for i := 0; i < len(e); i++ {
			fmt.Print(string(c[i%len(c)] ^ e[i]))
		}
		return
	}

	f := make(chan bool)
	for i := 0; i < len(b); i++ {
		go func(i int) {
			f <- b[i] == a[len(a)-1-i%len(a)]^os.Args[1][i]
		}(i)
	}
	r := true
	for i := 0; i < len(b); i++ {
		r = r && <-f
	}
	if r {
		for i := 0; i < len(d); i++ {
			fmt.Print(string(c[i%len(c)] ^ d[i]))
		}
	} else {
		for i := 0; i < len(e); i++ {
			fmt.Print(string(c[i%len(c)] ^ e[i]))
		}
	}
}
