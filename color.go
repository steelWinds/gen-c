package genc

import (
	"github.com/lucasb-eyer/go-colorful"
	mclusters "github.com/muesli/clusters"
	"github.com/pkg/errors"
)

type ColorModelRGB = uint8
type ColorModelLab = float64

type ColorModels interface {
	~ColorModelRGB | ~ColorModelLab
}

func GenColors[T ColorModels](params ColorPaletteParams) (colors [][]T, err error) {
	var t T

	var clusters mclusters.Clusters

	if clusters, err = GetColorPalette(params); err != nil {
		err = errors.Wrap(err, "while gen palette throw error")

		return
	}

	colors = make([][]T, 0, len(clusters))

	for _, cItem := range clusters {
		center := cItem.Center
		color := colorful.Lab(center[0], center[1], center[2])

		var colorItem []T

		switch any(t).(type) {
		case ColorModelLab:
			l, a, b := color.Lab()

			colorItem = []T{T(l), T(a), T(b)}
		case ColorModelRGB:
			r, g, b := color.Clamped().RGB255()

			colorItem = []T{T(r), T(g), T(b)}
		}

		colors = append(colors, colorItem)
	}

	return
}
