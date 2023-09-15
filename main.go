package main

import "fmt"

func main() {
	var somoni int
	var rubl int = 10
	var dollar int = 1

	fmt.Scanf("%v", &somoni)
	println("Курс сомони в рублях ", somoni*rubl)
	println("Курс сомони в долларах ", somoni*dollar)
}
