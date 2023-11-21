package handlers

import (
	"os"
	"testing"
)

func TestConvertTimeToInt(t *testing.T) {
	if t2i(700) != 0 {
		t.Error("700 convert to 0 error")
	}

	if t2i(730) != 1 {
		t.Error("730 convert to 1 error")
	}

	if t2i(1200) != 10 {
		t.Error("1200 convert to 10 error")
	}
}

func TestGenSvg(t *testing.T) {
	rooms := make([]roomOccupiedArray, 3)
	rooms[0] = roomOccupiedArray{1, "第一会议室",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0, 0}}
	rooms[1] = roomOccupiedArray{2, "第二会议室",
		[]int{0, 0, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 0, 5, 5, 0, 11, 11, 11}}
	rooms[2] = roomOccupiedArray{3, "C206",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 21, 21, 21, 0, 0, 0, 0, 0, 0, 0, 31, 31, 31, 31, 31}}

	buffer := GenSvg(rooms)
	//write to file
	f, err := os.Create("test.svg")
	if err != nil {
		t.Errorf("create file error: %v", err)
		return
	}
	defer f.Close()

	n, err := f.WriteString(buffer.String())
	if err != nil {
		t.Errorf("write file error: %v", err)
		return
	}
	t.Logf("write %d bytes\n", n)
}
