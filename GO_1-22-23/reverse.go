package main

import "fmt";

func main() {
	var num, a, reverse int ;
	fmt.Println("Enter a number")
	fmt.Scanf("%d", & num)

	for num > 0 {
		a = num % 10;
		reverse = reverse * 10 + a;
		num = num /10
	}
	fmt.Println("Reverse of the number is: " , reverse)
}