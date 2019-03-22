/*
Find colors closest to the given color from the predefined set of colors.
*/
package closest_colors

import (
	"fmt"
	"math"
	"sort"
)

type RGBColor interface {
	RGB() (r, g, b uint8)
}

type ColorDistance struct {
	color    RGBColor
	distance float64
}

func getDistance(color1, color2 RGBColor) float64 {
	red1, green1, blue1 := color1.RGB()
	red2, green2, blue2 := color2.RGB()
	redDiff := float64(red1) - float64(red2)
	greenDiff := float64(green1) - float64(green2)
	blueDiff := float64(blue1) - float64(blue2)
	return math.Sqrt(redDiff*redDiff + greenDiff*greenDiff + blueDiff*blueDiff)
}

func FindClosestRGBColors(color RGBColor, number int, colors []RGBColor) ([]ColorDistance, error) {
	if number <= 0 || number > len(colors) {
		return nil, fmt.Errorf("FindClosestRGBColors: number must be between 1 and colors size, got %d", number)
	}

	colorDistances := make([]ColorDistance, 0, len(colors))
	for _, c := range colors {
		colorDistances = append(colorDistances, ColorDistance{c, getDistance(color, c)})
	}
	sort.Slice(colorDistances, func(i, j int) bool {
		return colorDistances[i].distance < colorDistances[j].distance
	})

	return colorDistances[:number], nil
}
