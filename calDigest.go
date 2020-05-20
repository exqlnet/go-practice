package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	// 选择摘要方法
	flag.Parse()
	t := flag.Arg(0)
	supports := [...]string{"sha256", "sha512", "sha384"}
	defaultType := "sha256"
	for i, s := range supports {
		if t == s {
			break
		}
		if i == len(supports)-1 {
			t = defaultType
		}
	}

	fmt.Printf("using type %s\n", t)

	var in string
	_, _ = fmt.Scanln(&in)

	switch t {
	case "sha256":
		c := sha256.Sum256([]byte(in))
		fmt.Printf("%x", c)
	case "sha384":
		c := sha512.Sum384([]byte(in))
		fmt.Printf("%x", c)
	case "sha512":
		c := sha512.Sum512([]byte(in))
		fmt.Printf("%x", c)
	}

}
