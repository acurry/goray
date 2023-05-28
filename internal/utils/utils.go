package utils

func LinSpace(lo, hi float64, num int) []float64 {
	result := make([]float64, 1, num)
	result[0] = lo
	diff := (hi - lo) / float64(num)
	curr := lo
	for i := 0; i < num-1; i++ {
		curr += diff
		result = append(result, curr)
	}
	result = append(result, hi)

	return result
}

func Zeroes(length, width, height int) [][][]float64 {
	rows := make([][][]float64, length)
	for r := 0; r < length; r++ {
		rows[r] = make([][]float64, width)
		for c := 0; c < width; c++ {
			rows[r][c] = make([]float64, height)
			for d := 0; d < height; d++ {
				rows[r][c][d] = 0
			}
		}
	}

	return rows
}
