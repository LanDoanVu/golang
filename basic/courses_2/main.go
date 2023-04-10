package main

import "fmt"

func main() {
	// var laBai string = "A Bich"

	// laBai := "A Bich"

	laBai := taoBaiMoi()

	fmt.Println("Goi Ham Tao Bao Moi", laBai)

	// array danh sách có số phần tử cố định
	var arrayA [5]string
	arrayA[0] = "J"
	arrayA[1] = "Q"
	arrayA[2] = "K"
	arrayA[3] = "A"
	arrayA[4] = "2"

	fmt.Println("Array A", arrayA)

	// slice danh sách có thể mở rộng được
	// Static Slice type string
	sliceA := make([]string, 3)
	// sliceA := []string{"J", "Q", "K"}
	sliceA[0] = "J"
	sliceA[1] = "Q"
	sliceA[2] = "K"

	//Them 1 phan tu vao slice
	sliceA = append(sliceA, "A", "2")
	fmt.Println("Slice A", sliceA)

	for index, laBai := range sliceA {
		fmt.Println("La bai thu", index, laBai)
	}
}

func taoBaiMoi() string {
	return "K co"
}
