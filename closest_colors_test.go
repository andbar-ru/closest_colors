package closest_colors

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

const (
	EPSILON = 0.01
)

type rgb struct {
	Red   uint8 `json:"r"`
	Green uint8 `json:"g"`
	Blue  uint8 `json:"b"`
}

type color struct {
	Id   int    `json:"colorId"`
	Rgb  rgb    `json:"rgb"`
	Hex  string `json:"hexString"`
	Name string `json:"name"`
}

func (rgb rgb) RGB() (uint8, uint8, uint8) {
	return rgb.Red, rgb.Green, rgb.Blue
}

func (c color) RGB() (uint8, uint8, uint8) {
	return c.Rgb.Red, c.Rgb.Green, c.Rgb.Blue
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func TestFindClosestRGBColors(t *testing.T) {
	jsonFile, err := os.Open("./testdata/term_colors.json")
	check(err)
	defer jsonFile.Close()

	var colors []color
	err = json.NewDecoder(jsonFile).Decode(&colors)
	check(err)

	// Convert to slice of interfaces.
	rgbColors := make([]RGBColor, len(colors))
	for i, c := range colors {
		rgbColors[i] = c
	}

	// Test single result.
	results, err := FindClosestRGBColors(&rgb{173, 43, 82}, 1, rgbColors)
	if err != nil {
		t.Error(err)
	} else if len(results) != 1 {
		t.Error("Expected 1, got", len(results))
	} else {
		result := results[0]
		resultColor := result.color.(color)
		if resultColor.Id != 125 || result.distance-44.97 > EPSILON {
			t.Errorf("Expected 125 and ~44.97, got %d and %.2f", resultColor.Id, result.distance)
		}
	}

	// Test multiple results.
	results, err = FindClosestRGBColors(&rgb{99, 97, 25}, 5, rgbColors)
	if err != nil {
		t.Error(err)
	} else if len(results) != 5 {
		t.Error("Expected 5, got", len(results))
	} else {
		expectedIds := [5]int{58, 94, 64, 3, 100}
		expectedDistances := [5]float64{25.40, 43.87, 45.66, 49.26, 58.01}
		for i, result := range results {
			resultColor := result.color.(color)
			if resultColor.Id != expectedIds[i] || result.distance-expectedDistances[i] > EPSILON {
				t.Errorf("result %d: expected id=%d and distance~%v, got %d and %.2f", i, expectedIds[i], expectedDistances[i], resultColor.Id, result.distance)
			}
		}
	}

	// Test errors
	results, err = FindClosestRGBColors(&rgb{173, 43, 82}, 0, rgbColors)
	if err == nil {
		t.Error("Expected not nil error")
	}
	results, err = FindClosestRGBColors(&rgb{173, 43, 82}, 1000, rgbColors)
	if err == nil {
		t.Error("Expected not nil error")
	}
}
