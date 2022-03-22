package area_test

import (
	. "go-tests/area"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 12}
		expectedArea := float64(120)
		currentArea := rectangle.Area()

		if expectedArea != currentArea {
			t.Fatalf("Expected %0.2f but received %0.2f", expectedArea, currentArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		expectedArea := float64(math.Pi * 100)
		currentArea := circle.Area()

		if expectedArea != currentArea {
			t.Fatalf("Expected %0.2f but received %0.2f", expectedArea, currentArea)
		}
	})
}
