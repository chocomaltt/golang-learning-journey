package task

func CelsiusConvertion(degree float32) [2]float32 {
	fahrenheit := (degree * 9 / 5) + 32
	kelvin := degree + 273.15

	results := [...]float32{fahrenheit, kelvin}
	return results
}
