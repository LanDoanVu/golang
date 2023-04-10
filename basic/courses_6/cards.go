package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (n nhieuLaBai) arrayToString() string {
	return strings.Join(n, ",")
}

func (n nhieuLaBai) saveFile(tenFile string) error {
	data := []byte(n.arrayToString())
	return ioutil.WriteFile(tenFile, data, 0666)
}

func readFile(tenFile string) nhieuLaBai {
	data, err := ioutil.ReadFile(tenFile)

	if err != nil {
		log.Printf("Read File Err %v", err)
		os.Exit(1)
	}

	ctxFile := string(data)

	ctxFileArrayToString := strings.Split(ctxFile, ",")

	convertFile := nhieuLaBai(ctxFileArrayToString)

	return convertFile
}

func (n nhieuLaBai) xaoBai() {
	rand.Seed(time.Now().UnixNano())
	for index := range n {
		numberRandom := rand.Intn(len(n))

		// Doi cho n[index] n[numberRandom]
		n[index], n[numberRandom] = n[numberRandom], n[index]
	}
}
