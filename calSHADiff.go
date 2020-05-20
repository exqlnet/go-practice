// 计算两个字符串经过sha256 hash过后不同的比特位数
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("abcd"))
	c2 := sha256.Sum256([]byte("efgh"))

	count := 0
	for i, _ := range c1 {
		tmp := c1[i] ^ c2[i]
		count += getNotZeroCount(tmp)
	}
	fmt.Printf("%x\n%x\n%d", c1, c2, count)

}

func getNotZeroCount(n byte) int {
	count := 0
	for n != 0 {
		if n&1 > 0 {
			count++
		}
		n >>= 1
	}
	return count
}
