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
