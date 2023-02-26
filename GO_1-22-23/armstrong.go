package main

import "fmt";

func main() {
	var num, r, sum, temp int = 1, 0, 0, 0;
	fmt.Println("Enter a number");
	fmt.Scanf("%d", &num);
	temp = num;
	for num!= 0 {
		r = num % 10;
		sum += r * r * r;
		num = num /10;
	}
	if temp == sum {
		fmt.Println("Armstrong number");
	} else {
		fmt.Println("Not an Armstrong number");
	}

}

