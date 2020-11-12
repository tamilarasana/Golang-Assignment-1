//Using  Function and Struct

package main

import "fmt"


// type User struct{
// 	Name string
// 	Phone int
// 	Age int
	
// }

type Address struct {
	User
	City string
	State string
	Email string
}

type User struct{
	Name string
	Phone int
	Age int
	City string
	State string
	Email string

}


func main(){
	var details User

	details.Name  = "Tamilarasan"
	details.Phone = 9787973192
	details.Age   = 24
	details.City  = "Krishnagiri" 
	details.State = "TamilNadu"
	details.Email = "tamil@gmail.com"

	printUser(details)
}

func printUser(userDetails User){

		fmt.Println(userDetails.Name)
		fmt.Println(userDetails.Phone)
		fmt.Println(userDetails.Age)
		fmt.Println(userDetails.City)
		fmt.Println(userDetails.State)
		fmt.Println(userDetails.Email)
}

/*package main

import "fmt"


type User struct{
	Name string
	Phone int
	Age int
	City string
	
}


func main(){

	//Creating variable

	details := User{"Tamilarasan", 9876543210,24, "Krishnagiri"}

	fmt.Println(details)
	fmt.Printf("%v\n",details)
	fmt.Printf("%+v\n",details)

    //Another Way

	var tamil User   //or  --> var tamil = new(User)
	tamil.Name = "Tamilarasan"
	tamil.Phone = 9876543210
	tamil.Age = 24
	tamil.City = "Krishnagiri" 

	fmt.Println(tamil)
	fmt.Printf("%v\n",tamil)
	fmt.Printf("%+v\n",tamil)

	 //Another way


	var data = &User{"Tamilarsan", 98765432123,24, "Krishnagiri"}

	fmt.Println(data)
	fmt.Println(&data) // store the data address
    fmt.Println(*data)// display the  address data
    fmt.Printf("%+v\n", tamil)
	

}*/

/*package main

import "fmt"


type User struct{
	Name string
	Phone int
	Age int
	
}

type Address struct {
	User
	City string
	State string
	Email string
}


func main(){

	Internal Struct

	a := Address{
		User: User{Name: "tamil", Phone: 987654321, Age: 24},
		City :"Hosur",
		State : "TamilNadu",
		Email :"tamil@gmail.com",

	}


	a := Address{}
	a.Name = "Tamilarasan"
	a.Phone =9787973192
	a.Age = 24
	a.City = "Hosur"
	a.State = "TamilNadu"
	a.Email = "tamil@gmail.com"


	fmt.Println(a)
}


Output

    {{Tamilarasan 9787973192 24} Hosur TamilNadu tamil@gmail.com}*/





