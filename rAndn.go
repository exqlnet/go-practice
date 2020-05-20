// 展示换行和回车字符的区别
// \r 回车 把打印的针头移动到左边届
// \n 换行
// 各操作系统文件行结尾：
// unix: \n
// mac: \r
// windows: \n\r
package main

import "fmt"

func main() {
	fmt.Printf("开始换行\n")
	fmt.Printf("开始换行\n")
	fmt.Printf("开始回车\r")
	fmt.Printf("回车后")

}
