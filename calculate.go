package main

import "math"

func getOrientation(rM []float64) (values []float64) {
	values = make([]float64, 3)
	if len(rM) == 9 {
		values[0] = float64(math.Atan2(rM[1], rM[4]))
		values[1] = float64(math.Asin(-rM[7]))
		values[2] = float64(math.Atan2(-rM[6], rM[8]))
		return values
	}
	values[0] = float64(math.Atan2(rM[1], rM[5]))
	values[1] = float64(math.Asin(-rM[9]))
	values[2] = float64(math.Atan2(-rM[8], rM[10]))
	return values
}
