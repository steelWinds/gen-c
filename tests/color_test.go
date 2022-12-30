package tests

import (
	"testing"

	"github.com/steelWinds/genc"
	"github.com/stretchr/testify/assert"
)

func TestGenColorRGB(t *testing.T) {
	const PartsAmount = 10
	const RGBModelLen = 3

	params := genc.ColorPaletteParams{
		StartLight:  -1,
		EndLight:    1,
		PartsAmount: PartsAmount,
	}

	colors, err := genc.GenColors[genc.ColorModelRGB](params)

	assert.Nil(t, err)

	assert.Len(t, colors, PartsAmount)

	for _, colorItem := range colors {
		assert.Len(t, colorItem, RGBModelLen)

		for _, n := range colorItem {
			assert.IsType(t, *new(genc.ColorModelRGB), n)
		}
	}
}

func TestGenColorLab(t *testing.T) {
	const PartsAmount = 10
	const LabModelLen = 3

	params := genc.ColorPaletteParams{
		StartLight:  -1,
		EndLight:    1,
		PartsAmount: PartsAmount,
	}

	colors, err := genc.GenColors[genc.ColorModelLab](params)

	assert.Nil(t, err)

	assert.Len(t, colors, PartsAmount)

	for _, colorItem := range colors {
		assert.Len(t, colorItem, LabModelLen)

		for _, n := range colorItem {
			assert.IsType(t, *new(genc.ColorModelLab), n)
		}
	}
}
