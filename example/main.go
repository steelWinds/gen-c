package main

import (
	"fmt"

	"github.com/steelWinds/gen-c/internal/gen"
	"github.com/steelWinds/gen-c/pkg"
)

func main() {
	params := gen.ColorPaletteParams{
		StartLight: -1,
		EndLight: 1,
		PartsAmount: 3,
	}

	colorsLab, err := pkg.GenColors[pkg.ColorModelLab](params)
	colorsRGB, err := pkg.GenColors[pkg.ColorModelRGB](params)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%v model colors: %v \n", "RGB", colorsRGB)
	fmt.Printf("%v model colors: %v", "Lab", colorsLab)
}
