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

// ヘルパー関数を作成していく
func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
