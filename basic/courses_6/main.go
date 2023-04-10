package main

import "fmt"

func main() {

	baiTrenTay := themLaBai()

	fmt.Println("Bai Hien Co")
	baiTrenTay.in()

	fmt.Println("Tron Bai")
	baiTrenTay.xaoBai()
	baiTrenTay.in()
}
