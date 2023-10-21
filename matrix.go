package main

import "math"

func getRotationMatrix(rM []float64, iM []float64, gravity []float64, geomagnetic []float64) bool {
	Ax := gravity[0]
	Ay := gravity[1]
	Az := gravity[2]
	normsqA := Ax*Ax + Ay*Ay + Az*Az
	g := 9.81
	freeFallGravitySquared := 0.01 * g * g
	if normsqA < freeFallGravitySquared {
		return false
	}
	Ex := geomagnetic[0]
	Ey := geomagnetic[1]
	Ez := geomagnetic[2]
	Hx := Ey*Az - Ez*Ay
	Hy := Ez*Ax - Ex*Az
	Hz := Ex*Ay - Ey*Ax
	normH := math.Sqrt(Hx*Hx + Hy*Hy + Hz*Hz)
	if normH < 0.1 {
		return false
	}
	invH := 1.0 / normH
	Hx *= invH
	Hy *= invH
	Hz *= invH
	invA := 1.0 / math.Sqrt(Ax*Ax+Ay*Ay+Az*Az)
	Ax *= invA
	Ay *= invA
	Az *= invA
	Mx := Ay*Hz - Az*Hy
	My := Az*Hx - Ax*Hz
	Mz := Ax*Hy - Ay*Hx
	if len(rM) > 0 {
		if len(rM) == 9 {
			rM[0] = Hx
			rM[1] = Hy
			rM[2] = Hz
			rM[3] = Mx
			rM[4] = My
			rM[5] = Mz
			rM[6] = Ax
			rM[7] = Ay
			rM[8] = Az
		} else if len(rM) == 16 {
			rM[0] = Hx
			rM[1] = Hy
			rM[2] = Hz
			rM[3] = 0
			rM[4] = Mx
			rM[5] = My
			rM[6] = Mz
			rM[7] = 0
			rM[8] = Ax
			rM[9] = Ay
			rM[10] = Az
			rM[11] = 0
			rM[12] = 0
			rM[13] = 0
			rM[14] = 0
			rM[15] = 1
		}
	}
	if len(iM) > 0 {
		invE := 1.0 / math.Sqrt(Ex*Ex+Ey*Ey+Ez*Ez)
		c := (Ex*Mx + Ey*My + Ez*Mz) * invE
		s := (Ex*Ax + Ey*Ay + Ez*Az) * invE
		if len(iM) == 9 {
			iM[0] = 1
			iM[1] = 0
			iM[2] = 0
			iM[3] = 0
			iM[4] = c
			iM[5] = s
			iM[6] = 0
			iM[7] = -s
			iM[8] = c
		} else if len(iM) == 16 {
			iM[0] = 1
			iM[1] = 0
			iM[2] = 0
			iM[3] = 0
			iM[4] = c
			iM[5] = s
			iM[6] = 0
			iM[7] = -s
			iM[8] = c
			iM[9] = 0
			iM[10] = 0
			iM[11] = 0
			iM[12] = 0
			iM[13] = 0
			iM[14] = 0
			iM[15] = 1
		}
	}
	return true
}
