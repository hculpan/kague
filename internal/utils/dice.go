package utils

import "math/rand/v2"

// GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	return rand.IntN(num)
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	return GetRandomInt(num) + 1
}

func GetRandomBetween(low int, high int) int {
	return GetRandomInt(high+1-low) + low
}
