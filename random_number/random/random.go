package random

import "math/rand"
import "com.shreyash/random-number/random/output"

func RandomNumber() int {
	outptut.OutputText("Random number generated")
	return rand.Intn(100)
}