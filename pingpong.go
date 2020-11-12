package main

import "fmt"

func main()  {

	fmt.Println("Ping And Pong Program")
	
	for num := 1; num<=50;  num++ {

		if num %15 == 0{
			fmt.Println("PingPong")
		}else if num % 3 == 0{
			fmt.Println("ping")
		}else if  num % 5 == 0{
			fmt.Println("Pong")
		}else{
			fmt.Println(num)
		}
		
	}
	
}