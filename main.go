package main

import "tests/exercise"

func main() {

	w, h, a, err := exercise.GetCharacter()
	if err != nil {
		return
	}
	dayLimit := exercise.CountRequiredCalories(w, h, a)
	expenditure := exercise.Expenditure(10, w, 5000)
	exercise.CheckBmr(dayLimit, expenditure)
}
