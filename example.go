package main

import "log"

var (
	/*
		MAG and GRAV are each an individual
		slice of values representing a magnetometer and
		accelerometer event respectively.
	*/
	MAGNETOMETER_EVENT = []float64{}
	GRAVITY_EVENT      = []float64{}
)

func Example() {
	rotationMatrix := make([]float64, 9)
	inclinationMatrix := make([]float64, 9)
	if getRotationMatrix(rotationMatrix, inclinationMatrix, GRAVITY_EVENT, MAGNETOMETER_EVENT) {
		orientation := getOrientation(rotationMatrix)
		log.Println(orientation)
	}
}
