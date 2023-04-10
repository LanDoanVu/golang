package main

import (
	"os"
	"testing"
)

func TestThemLaBai(t *testing.T) {
	boBai := themLaBai()

	if len(boBai) != 52 {
		t.Errorf("Something err %v", len(boBai))
	}
}

func TestChiaBai(t *testing.T) {
	boBai := themLaBai()

	tu1, tu2, tu3, tu4 := chiaBai(boBai, 13)

	if len(tu1) != 13 {
		t.Errorf("Something err %v", len(tu1))
	}

	if len(tu2) != 13 {
		t.Errorf("Something err %v", len(tu2))
	}

	if len(tu3) != 13 {
		t.Errorf("Something err %v", len(tu3))
	}

	if len(tu4) != 13 {
		t.Errorf("Something err %v", len(tu4))
	}
}

func TestSaveFileReadFile(t *testing.T) {
	boBai := nhieuLaBai{"Joker"}
	err := boBai.saveFile("_testCards")

	if err != nil {
		t.Errorf("something err %v", err)
	}

	boBai, err = readFile("_testCards")

	if err != nil {
		t.Errorf("Something err %v", err)
	}

	if len(boBai) != 1 {
		t.Errorf("Something err %v", len(boBai))
	}

	os.Remove("_testCards")
}
