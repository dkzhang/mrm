package handlers

import "testing"

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

//#FF0000 - Red
//#00FFFF - Aqua
//#800080 - Purple
//#FFFF00 - Yellow
//#0000FF - Blue
//#D2691E - Chocolate
//#008000 - Green
//#FF00FF - Fuchsia
//#A52A2A - Brown
//#FFA500 - Orange
//#008080 - Teal
//#6B8E23 - Olive Drab
//#000000 - Black
//#7FFFD4 - Aquamarine
//#808000 - Olive
//#800000 - Maroon
