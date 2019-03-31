package structsmethodsinterfaces

func perimeter(width, height float64) float64 {
	if width < 0 || height < 0 {
		return 0
	}
	return 2 * (width + height)
}

func area(width, height float64) float64 {
	if width < 0 || height < 0 {
		return 0
	}
	return width * height
}
