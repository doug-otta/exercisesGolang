package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	testsPerimeter := []struct {
		name           string
		shape          Shape
		totalPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 12.0, width: 6.0}, totalPerimeter: 36.0},
		{name: "Circle", shape: Circle{radius: 10.0}, totalPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{height: 12.0, width: 6.0}, totalPerimeter: 18.0},
	}

	for _, tt := range testsPerimeter {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Perimeter()
			expected := tt.totalPerimeter
			if result != expected {
				t.Errorf("%#v RESULT: %.2v, EXPECTED: %.2v", tt.shape, result, tt.totalPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		name      string
		shape     Shape
		totalArea float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 12, width: 6}, totalArea: 72},
		{name: "Circle", shape: Circle{radius: 10}, totalArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{height: 12, width: 6}, totalArea: 36},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.totalArea {
				t.Errorf("%#v RESULT: %.2v, EXPECTED: %.2v", tt.shape, result, tt.totalArea)
			}
		})
	}
}
