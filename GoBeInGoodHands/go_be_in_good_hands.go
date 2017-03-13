package main

import (
	"math"
)

// Oh no! This is a bogus function!
func wrong_pow(value float64, to int64) float64 {
	if to == 0 {
		return 1
	}
	if to == 1 {
		return value
	}
	answer := wrong_pow(value, to/2)
	return answer * answer
}

func pow(value float64, to float64) float64 {
	if to == 0 {
		return 1
	}
	if to == 1 {
		return value
	}
	halfTo := to / 2
	return pow(value, math.Floor(halfTo)) * pow(value, math.Ceil(halfTo))
}

func square(value float64) float64 {
	return pow(value, 2)
}

func main() {

}
