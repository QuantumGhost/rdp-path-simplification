package rdp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	points := []Point{
		Point{0, 0},
		Point{1, 2},
		Point{2, 7},
		Point{3, 1},
		Point{4, 8},
		Point{5, 2},
		Point{6, 8},
		Point{7, 3},
		Point{8, 3},
		Point{9, 0},
	}

	t.Run("Threshold=0", func(t *testing.T) {
		if len(SimplifyPath(points, 0)) != 10 {
			t.Error("simplified path should have all points")
		}
	})

	t.Run("Threshold=2", func(t *testing.T) {
		if len(SimplifyPath(points, 2)) != 7 {
			t.Error("simplified path should only have 7 points")
		}
	})

	t.Run("Threshold=5", func(t *testing.T) {
		if len(SimplifyPath(points, 100)) != 2 {
			t.Error("simplified path should only have two points")
		}
	})
}

func TestSeekMostDistantPoint(t *testing.T) {
	l := Line{Start: Point{0, 0}, End: Point{10, 0}}
	points := []Point{
		Point{13, 13},
		Point{1, 15},
		Point{1, 1},
		Point{3, 6},
	}

	idx, maxDist := seekMostDistantPoint(l, points)

	if idx != 1 {
		t.Error("failed to find most distant point away from a line")
	}

	if maxDist != 15 {
		t.Error("maximum distance is incorrect")
	}
}

func TestInfiniteRecursion(t *testing.T) {
	type LatLng struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	}

	file, err := os.Open("./testdata/infinite-recursion-1.json")
	require.NoError(t, err)
	var latLngs []LatLng
	err = json.NewDecoder(file).Decode(&latLngs)
	require.NoError(t, err)
	rdpPoints := make([]Point, 0, len(latLngs))
	for _, latLng := range latLngs {
		rdpPoints = append(rdpPoints, Point{
			X: latLng.Longitude,
			Y: latLng.Latitude,
		})
	}
	simplified := SimplifyPath(rdpPoints, 5e-05)
	assert.NotEmpty(t, simplified)
}
