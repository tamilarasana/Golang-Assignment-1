    
//Channels is a Conneting pipe like two goroutine  conneting


package main

import "fmt"

func sentValue(details chan string){

	details <- "Tamilarasan"
	details <- "Annamalai"
	details <- "Krishnagiri"
}

func main (){

	values := make(chan string)
    
	go sentValue(values)

	value :=  <- values

	defer fmt.Println(value)
		  fmt.Println(<-values)
		  fmt.Println(<-values)
}


// package main

// import "fmt"

// func main(){

// 	myChannel := make(chan string)

// 	go func(){
// 		myChannel <- "Tamilarasan"
// 		myChannel <- "Kloudone"
// 	}()

// 	text := <-myChannel

// 	fmt.Println(text)
// 	fmt.Println(<-myChannel)
// }