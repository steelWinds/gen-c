package main

import (
	"fmt"

	"github.com/steelWinds/genc"
)

func main() {
	params := genc.ColorPaletteParams{
		StartLight:  -1,
		EndLight:    1,
		PartsAmount: 3,
	}

	var colorsLab [][]genc.ColorModelLab
	var colorsRGB [][]genc.ColorModelRGB
	var err error

	if colorsLab, err = genc.GenColors[genc.ColorModelLab](params); err != nil {
		panic(err)
	}

	if colorsRGB, err = genc.GenColors[genc.ColorModelRGB](params); err != nil {
		panic(err)
	}

	fmt.Printf("%v model colors: %v \n", "RGB", colorsRGB)
	fmt.Printf("%v model colors: %v", "Lab", colorsLab)
}
