package utils

func TransformCurrentState(currentState string) float64 {
	var state float64
	if currentState == "transmitting" {
		state = 1
	} else {
		state = 0
	}
	return state
}
