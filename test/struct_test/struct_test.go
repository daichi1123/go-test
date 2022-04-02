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
// テーブル駆動テストの方法
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
