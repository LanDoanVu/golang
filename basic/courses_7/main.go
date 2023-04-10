package main

import (
	"fmt"
	"os"
)

func main() {

	baiTrenTay := themLaBai()

	fmt.Println("Bai Hien Co")
	baiTrenTay.in()

	baiTrenTay.saveFile("Cards")
	baiTrenTay = append(baiTrenTay, "Joker")
	fmt.Println("Doc Tu File")
	baiTrenTay.in()

	baiTrenTay, err := readFile("Cards")
	if err != nil {
		os.Exit(1)
	}

	baiTrenTay.in()
}
