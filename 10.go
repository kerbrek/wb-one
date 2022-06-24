package main

import "fmt"

// Дана последовательность температурных колебаний:
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножествах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func groupByRange(fluctuations []float64) map[int][]float64 {
	res := make(map[int][]float64)

	for _, f := range fluctuations {
		step := (int(f) / 10) * 10
		res[step] = append(res[step], f)
	}

	return res
}

func groupByRangeUniq(fluctuations []float64) map[int]map[float64]struct{} {
	resUniq := make(map[int]map[float64]struct{})

	for _, f := range fluctuations {
		step := (int(f) / 10) * 10
		if _, ok := resUniq[step]; !ok {
			resUniq[step] = make(map[float64]struct{})
		}

		resUniq[step][f] = struct{}{}
	}

	return resUniq
}

func main() {
	fluctuations := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(groupByRange(fluctuations))

	fmt.Println(groupByRangeUniq(([]float64{-25.4, -27.0, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})))
}
