package main

func main() {
	// baiTrenTay := nhieuLaBai{"J", "Q", "K"}

	baiTrenTay := themLaBai()

	// for index, laBai := range baiTrenTay {
	// 	fmt.Println(index, laBai)
	// }

	baiTrenTay.in()
}

func themLaBai() nhieuLaBai {
	return nhieuLaBai{"A", "2"}
}
