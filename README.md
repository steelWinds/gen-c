[![Go Reference](https://pkg.go.dev/badge/github.com/steelWinds/gen-c.svg)](https://pkg.go.dev/github.com/steelWinds/gen-c)
[![Go Report Card](https://goreportcard.com/badge/github.com/steelWinds/gen-c)](https://goreportcard.com/report/github.com/steelWinds/gen-c)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/steelWinds/gen-c/test-actions.yml?label=tests)

# genc
Simple console-aplication for generate colors in **RGB** and **Lab** models.

## Getting Started
Try

```go
params := genc.ColorPaletteParams{
    StartLight: -1,
    EndLight: 1,
    PartsAmount: 3,
}

var err error
var colorsLab [][]genc.ColorModelLab
var colorsRGB [][]genc.ColorModelRGB

// In Lab model
colorsLab, err = genc.GenColors[genc.ColorModelLab](params)

// In RGB model
colorsRGB, err = genc.GenColors[genc.ColorModelRGB](params)
```

## License
This project under Mozilla Public License Version 2.0 license, see more in LICENSE file

## Author

[@steelWinds](https://github.com/steelWinds)
