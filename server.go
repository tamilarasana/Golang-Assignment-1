package main

import (
	"fmt"
	"net/http"
)

func myName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hey hi this is Tamilarasan")
}


func helloEveryone(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Every  one")
}


func main(){

	http.HandleFunc("/",  myName)
	http.HandleFunc("/tamil", helloEveryone)
	http.ListenAndServe(":9000",nil)


}

// func main(){




// 	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w,"Hey hi this is Tamilarasan")
// })

// 	http.HandleFunc("/tamil", func(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "No Love bro !! Single morratu Single")
// })
// http.ListenAndServe(":9000",nil)

