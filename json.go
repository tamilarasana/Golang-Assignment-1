package main


import (
	    "fmt"
		"encoding/json"
		)	

type Preson struct{
	Name string `json:"name"`
	Phone int   `json:"phone"`
	Age int     `json:"age"`
	Gender string `json:"gender"`
	City []string `json:"city"`
}

func main () {

	user := &Preson{Name :"Tamilarasan",Phone:9787973192,Age:24,Gender: "Male", City: []string{"Krishnagiri", "Hosur"}}
	data, _:= json.Marshal(user)
	fmt.Println(string(data))


	 user2 := `{"Name":"Tamilarasan","Phone":9787973192,"Age":24,"Gender":"Male","City":["Krishnagiri","Hosur"]}`
	  data2 := &Preson{}

	  json.Unmarshal([]byte(user2),data2)
	  fmt.Println(data2)
}