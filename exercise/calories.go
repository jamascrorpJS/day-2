package exercise

import "fmt"

const weightCof = 13.4
const heightCof = 4.8
const ageCof = 5.7
const bmr = 88.36

func GetCharacter() (weight, height, age float64, err error) {
	fmt.Println("Введите вес в кг, рост в см, возраст в годах!")
	_, err = fmt.Scan(&weight, &height, &age)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
		return
	}
	if weight <= 0 {
		fmt.Println("Вес не может быть <= 0!")
		return
	}
	return
}

// BMR = 88,36 + (13,4 × вес в кг) + (4,8 × рост в см) – (5,7 × возраст в годах).
func CountRequiredCalories(weight, height, age float64) float64 {
	dayLimit := bmr + (weightCof * weight) + (heightCof * height) - (ageCof * age)
	fmt.Printf("Вам необходимо %.2f калорий в день\n", dayLimit)
	return dayLimit
}

func CheckBmr(dayLimit, expenditure float64) {
	if dayLimit >= expenditure {
		fmt.Println("У вас недобор! Ешьте побольше!")
		return
	} else {
		fmt.Println("У вас перебор! Ешьте поменьше!")
		return
	}
}

func Expenditure(kilometers, weight, eatenKal float64) float64 {
	fmt.Println("Введите дистанцию которую вы пробежали и калории")
	fmt.Scan(&kilometers, &eatenKal)
	return eatenKal - (kilometers * weight)
}
