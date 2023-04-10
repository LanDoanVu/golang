package main

import "fmt"

func main() {

	baiTrenTay := themLaBai()

	fmt.Println("Cac Quan Bai Trong Bo Bai")
	baiTrenTay.in()

	baiDaChia, baiConLai := chiaBai(baiTrenTay, 9)

	fmt.Println("Bai Da Chia")
	baiDaChia.in()

	fmt.Println("Cac Quan Bai Con Lai")
	baiConLai.in()
}
