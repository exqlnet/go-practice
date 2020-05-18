// 消除相邻重复字符串，原地算法
// T(n) = O(n)
// S(n) = O(1)
package main

import "fmt"

func trimDuplicatedChar(strs []string)  []string {
	i := -1
	j := 0

	for ;j < len(strs); j++ {
		if i == -1 || strs[i] != strs[j] {
			i ++
			strs[i] = strs[j]
		} else if strs[i] == strs[j]{
			i --
		}
	}

	return strs[:i+1]
}

func main() {
	fmt.Printf("%v", trimDuplicatedChar([]string{"a", "a", "b", "b", "c"}))
}