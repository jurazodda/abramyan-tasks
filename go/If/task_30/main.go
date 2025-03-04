package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if (num >= 10 && num <= 99) && num % 2 == 0 {
		fmt.Println("Четное двузначное число")
	} else if (num >= 100 && num <= 999) && num % 2 == 1 {
		fmt.Println("Нечетное трехзначное число")
	}
}