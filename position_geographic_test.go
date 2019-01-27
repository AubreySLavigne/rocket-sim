package main

import (
	"math"
	"testing"
)

func TestGeographicToCartesian(t *testing.T) {
	tests := []struct {
		lat      radians
		long     radians
		altitude float64

		expects positionCartesian
	}{
		{lat: 0, long: 0, altitude: 0, expects: positionCartesian{x: 0, y: 0, z: 0}},
		{lat: 0, long: 0, altitude: 1000, expects: positionCartesian{x: 1000, y: 0, z: 0}},
		{lat: 0, long: 0.5 * math.Pi, altitude: 1000, expects: positionCartesian{x: 0, y: 1000, z: 0}},
		{lat: 0, long: math.Pi, altitude: 1000, expects: positionCartesian{x: -1000, y: 0, z: 0}},
	}

	for _, test := range tests {

		pos := positionGeographic{
			lat:      test.lat,
			long:     test.long,
			altitude: test.altitude,
		}
		if res := pos.toCartesian(); !res.equal(test.expects) {
			t.Errorf("Geographic %v incorrectly converted to Cartesian. Got %v, Expected %v", pos, res, test.expects)
		}
	}
}
