package structsmethodsinterfaces

func Perimeter(rectangle Rectangle) float64 {
	return rectangle.Length*2 + rectangle.Width*2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Length * rectangle.Width
}
