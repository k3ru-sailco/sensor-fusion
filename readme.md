# sensor-fusion

Create a virtual orientation event by fusing gravity and magnetometer sensor data together. This is a port of the respective methods from the [Android Device SensorManager](https://developer.android.com/reference/android/hardware/SensorManager) to Go for use with spoofed or synthetic events.

This replicates how an android device typically gets its 3 orientation angles.

The downside to the use of this method is that the orientation events are not as exact as combined sensor would be due to the fact that the magnetometer and gravity sensor are inherently noisey/inaccurate.

More information regarding this topic can be found on a [blog post](https://plaw.info/articles/sensorfusion/) made by Paul Lawitzki on Android sensor fusion, and applying a complimentrary filter to reduce or eradicate gyroscope drift and the aforementioned magnetometer and gravity drift. This post also served as a great reference while porting similar methods to Go.

## Usage
```go
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
```
