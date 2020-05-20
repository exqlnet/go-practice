package main

import (
	"fmt"
	"io"
)

func main() {
	var a []string
	a = make([]string, 1, 8)
	fmt.Printf("%T\n%t\n%d\n%d", a, a == nil, len(a), cap(a))
	//f(buf)
	//f(buf)
}

func f(out io.Writer) {

	fmt.Printf("%T\n%t\n", out, out == nil)
	//if out != nil {
	//	out.Write([]byte("done"))
	//}
}
