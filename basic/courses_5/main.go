package main

import "fmt"

func main() {

	baiTrenTay := themLaBai()

	// fmt.Println("Cac Quan Bai Trong Bo Bai")
	// baiTrenTay.in()

	// baiDaChia, baiConLai := chiaBai(baiTrenTay, 9)

	// fmt.Println("Bai Da Chia")
	// baiDaChia.in()

	// fmt.Println("Cac Quan Bai Con Lai")
	// baiConLai.in()

	baiTrenTay = append(baiTrenTay, "Joker")
	fmt.Println("Bai Tren Tay Sau Khi Da Them")
	baiTrenTay.in()

	fmt.Println("Convert Array To String and Save data to file")
	fmt.Println(baiTrenTay.arrayToString())
	baiTrenTay.saveFile("Cards")

	baiTrenTay = readFile("Cards")
	fmt.Println("Doc Du Lieu Tu File")
	baiTrenTay.in()

}
