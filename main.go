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
		PartsAmount: 15,
	}

	colors, err := pkg.GenColor[pkg.ColorModelLab](params)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(colors)
}
