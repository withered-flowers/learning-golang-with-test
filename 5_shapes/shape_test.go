package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// ? Using interface to determine the Shape and calculate the Area
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Calculating Area for Rectangles", func(t *testing.T) {
		// rectangle := Rectangle{12.0, 6.0}

		// got := rectangle.Area()
		// want := 72.0

		// if got != want {
		// 	t.Errorf("got %g want %g", got, want)
		// }

		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rectangle, want)
	})

	t.Run("Calculating Area for Circle", func(t *testing.T) {
		// circle := Circle{10.0}

		// got := circle.Area()
		// want := 314.1592653589793

		// if got != want {
		// 	t.Errorf("got %g want %g", got, want)
		// }

		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

// ? Now we will using the Table Driven Test
func TestArea2(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// ? Declaring Slice of Struct
		{
			name: "Rectangle",
			// ? the Struct
			shape: Rectangle{12, 6},
			// ? want
			hasArea: 72,
		},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	// ? Loop it
	for _, tt := range areaTest {

		// got := tt.shape.Area()

		// if got != tt.hasArea {
		// 	t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		// }

		// ! Test should be as verbose as possible
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
