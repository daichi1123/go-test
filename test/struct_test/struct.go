package structtest

type Rectangle struct {
	Width  float64
	Height float64
}

// テストの正解
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
