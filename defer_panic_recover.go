package main

import "fmt"

func main(){
	defer catchIssues("divide")  
	divide()
	fmt.Println("End of the Main")
}

func catchIssues(functionName string){
	rec := recover()
	if rec != nil{
		fmt.Println("somthing has gone  worng",functionName)
		fmt.Println("Reson:",rec)
	}
}

func divide(){

	num := 20
	den := 0
	fmt.Println("Enter a number")
	fmt.Scanf("%d",&den)

	if den == 0 {
		panic("Oops !!! Cannot Enter 0")
	}else{
		div := num/den
		fmt.Println(div)
	}
}




// package main

// import "fmt"

// func div(num1 int, num2 int){
// 	defer func(){

// 		if  i := recover(); i != nil{
// 			fmt.Printf("%v",i)
// 		} 
// 	}()

// 	divide := num1/num2
// 	fmt.Printf("Divide %vln",divide)
// }

// func main(){

// 	//fmt.Printf("Entert The Number")
// 	 num1 := 20
// 	 num2 :=10
// 	//fmt.Scanf("%d",num1)
// 	//fmt.Scanf("%d", num2)
// 	div(num1,num2)

// }