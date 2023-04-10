package main

import "fmt"

type nhieuLaBai []string

func themLaBai() nhieuLaBai {
	resutl := nhieuLaBai{}

	cacChatBai := []string{"co", "ro", "chuon", "bich"}
	cacQuanBai := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, chatBai := range cacChatBai {
		for _, quanBai := range cacQuanBai {
			bai := quanBai + " " + chatBai

			resutl = append(resutl, bai)
		}
	}

	return resutl
}

func chiaBai(n nhieuLaBai, sl int) (nhieuLaBai, nhieuLaBai) {
	return n[0:sl], n[sl:]
}

func (n nhieuLaBai) in() {
	fmt.Println(n)
}
