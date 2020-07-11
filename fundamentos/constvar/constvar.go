package main

import "math"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// Forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)

	println(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	println(c, d)

	var e, f = true, false
	println(e, f)

	g, h, i := 2, "epa!", false
	println(g, h, i)
}
