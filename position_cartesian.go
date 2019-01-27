package main

type positionCartesian struct {
	x float64
	y float64
	z float64
}

func (pc *positionCartesian) equal(other positionCartesian) bool {
	return pc.x == other.x && pc.y == other.y && pc.z == other.z
}
