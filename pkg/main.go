package main

import (
	"github.com/lucasb-eyer/go-colorful"
	mclusters "github.com/muesli/clusters"
	"github.com/steelWinds/gen-c/internal/gen"
)

type ColorModelRGB = uint8
type ColorModelLab = int8

type ColorModels interface {
	~ColorModelRGB | ~ColorModelLab
}

func genColor[T ColorModels](params gen.ColorPaletteParams, cModel T) (color string, err error) {
	var clusters mclusters.Clusters

	if clusters, err = gen.GetColorPalette(params); err != nil {
		return
	}

	colors := make([]T, len(clusters))

	for _, cItem := range clusters {
		center := cItem.Center
		color := colorful.Lab(center[0], center[1], center[2])

		

		colors = append(colors, color)
	}
}
