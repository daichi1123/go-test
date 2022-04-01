package structtest

import "testing"

func TestPerimeter(t *testing.T) {
	rectagle := Rectangle{10.0, 10.0}
	got := Perimeter(rectagle)
	want := 40.0

	if got != want {
		// .数字f ← これで小数点何桁かを指定して表示することができる
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
