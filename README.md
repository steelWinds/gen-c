[![Go Reference](https://pkg.go.dev/badge/github.com/steelWinds/gen-c.svg)](https://pkg.go.dev/github.com/steelWinds/gen-c)
[![Go Report Card](https://goreportcard.com/badge/github.com/steelWinds/gen-c)](https://goreportcard.com/report/github.com/steelWinds/gen-c)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/steelWinds/gen-c/test-actions.yml?label=tests)

# genc
Simple console-aplication for generate color palette.

## Getting Started

Try
```
params := genc.ColorPaletteParams{
    StartLight: -1,
    EndLight: 1,
    PartsAmount: 3,
}

colors, err := genc.GenColors(params)
```
