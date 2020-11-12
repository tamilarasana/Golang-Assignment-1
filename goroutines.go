package main

import(
	"fmt"
	"time"
	//"sync"
)

//var wg sync.WaitGroup

func sayHello(s string){
	for i :=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
	//wg.Done()
}

func main(){
	//wg.Add(1)
	go sayHello("hai")
	//wg.Add(1)
	go sayHello("Tamil")
	//wg.Add(1)
	go sayHello("bye")
	//wg.Wait()
	fmt.Println("done")
	time.Sleep(time.Second)
	fmt.Println("Success")
}


/*package main

import(
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func sayHello(s string){
	for i :=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main(){
	wg.Add(1)
	go sayHello("hai")
	wg.Add(1)
	go sayHello("Tamil")
	wg.Add(1)
	go sayHello("bye")
	wg.Wait()
	fmt.Println("done")
	time.Sleep(time.Second)
	fmt.Println("Success")
}*/