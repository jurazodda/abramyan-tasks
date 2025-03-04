package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	/*
		if num < 0{
			fmt.Print("Отрицательное ")
		} else {
			fmt.Print("Положительное ")
		}
		if num % 2 == 0 {
			fmt.Println("четное число")
		} else {
			fmt.Println("нечетное число")
		}
	*/

	if num < 0 && num % 2 == 0 {
		fmt.Println("Отрицательное четное число")
	} else if num == 0 {
		fmt.Println("Нулевое число")
	} else if num > 0 && num % 2 == 1 {
		fmt.Println("Положительное нечетное число")
	}
}
