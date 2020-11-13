package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // Game on; Toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball

	panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}





















// package main

// import "fmt"

// func main()  {

// 	fmt.Println("Ping And Pong Program")
	
// 	for num := 1; num<=50;  num++ {

// 		if num %15 == 0{
// 			fmt.Println("PingPong")
// 		}else if num % 3 == 0{
// 			fmt.Println("ping")
// 		}else if  num % 5 == 0{
// 			fmt.Println("Pong")
// 		}else{
// 			fmt.Println(num)
// 		}
		
// 	}
	
// }