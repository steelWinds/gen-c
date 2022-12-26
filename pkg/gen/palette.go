package gen

import (
	mclusters "github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func GetColorPalette(startL, endL float64, parts int) (clusters mclusters.Clusters, err error) {
	paletteObs := make(mclusters.Observations, 0)

	for l := float64(0); l < endL; l += 0.05 {
		for a := float64(-1); a < 1; a += 0.1 {
			for b := float64(-1); b < 1; b += 0.1 {
				paletteObs = append(paletteObs, mclusters.Coordinates{l, a, b})
			}
		}
	}

	km := kmeans.New()
	clusters, err = km.Partition(paletteObs, parts)

	if err != nil {
		return 
	}

	return
}