package main

import "math"

type positionGeographic struct {
	lat      radians
	long     radians
	altitude float64
}

// Assuming no rotation, a coordinate of (0, 0) will refer to the x-axis.
// Increasing along the Longitude moves the point in the y-plane
// (horizontal), while increasing along the latitude will move the point in
// the z-plane (vertical
//
// To handle floating-point precision errors, results are rounded to 10
// decimal places
func (pos *positionGeographic) toCartesian() positionCartesian {
	// TODO: Include the calculation for the z-vector. What's included here
	// only accounts for values in the x-y plane
	return positionCartesian{
		x: floatPrecision(math.Cos(float64(pos.long))*pos.altitude, 10),
		y: floatPrecision(math.Sin(float64(pos.long))*pos.altitude, 10),
		z: 0,
	}
}

// floatPrecision rounds the float to 'prec' digits of precision.
//
// For example, 3.1415926 rounded to 4 point of precision will result 3.1416
func floatPrecision(f float64, prec int) float64 {
	return math.Round(f*math.Pow10(prec)) / math.Pow10(prec)
}
