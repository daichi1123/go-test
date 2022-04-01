package structtest

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		// .数字f ← これで小数点何桁かを指定して表示することができる
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
