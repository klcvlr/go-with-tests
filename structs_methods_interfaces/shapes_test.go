package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	actual := rectangle.Perimeter()
	expected := 40.0

	if actual != expected {
		t.Errorf("Expected %2.f, but got %2.f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		expected := 72.0

		checkArea(t, rectangle, expected)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})

	t.Run("Circle and Rectangle with 'Table Driven Test", func(t *testing.T) {
		areaTests := []struct {
			name         string
			shape        Shape
			expectedArea float64
		}{
			{"Rectangle", Rectangle{12, 6}, 72.0},
			{"Circle", Circle{10}, 314.1592653589793},
			{"Triangle", Triangle{12, 6}, 36},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				actual := tt.shape.Area()
				if actual != tt.expectedArea {
					t.Errorf("Expected shape %#v to have area %g, but got %g", tt.shape, tt.expectedArea, actual)
				}
			})
		}
	})

}

func checkArea(t testing.TB, shape Shape, expected float64) {
	t.Helper()
	actual := shape.Area()

	if actual != expected {
		t.Errorf("Expected %g, but got %g", expected, actual)
	}
}
