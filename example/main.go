package main

import (
	"fmt"

	"github.com/steelWinds/gen-c/internal/gen"
	"github.com/steelWinds/gen-c/pkg"
)

func main() {
	params := gen.ColorPaletteParams{
		StartLight:  -1,
		EndLight:    1,
		PartsAmount: 3,
	}

	var colorsLab [][]pkg.ColorModelLab
	var colorsRGB [][]pkg.ColorModelRGB
	var err error

	if colorsLab, err = pkg.GenColors[pkg.ColorModelLab](params); err != nil {
		panic(err)
	}

	if colorsRGB, err = pkg.GenColors[pkg.ColorModelRGB](params); err != nil {
		panic(err)
	}

	fmt.Printf("%v model colors: %v \n", "RGB", colorsRGB)
	fmt.Printf("%v model colors: %v", "Lab", colorsLab)
}
