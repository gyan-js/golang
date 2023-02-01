package main

import "fmt"	

func main() {
	var num, product int = 0, 1;
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &num)

	for i := 0; i <= 10; i++ {
		product = num * i;
		fmt.Println(num, " * ", i, " = ", product);
	}
}