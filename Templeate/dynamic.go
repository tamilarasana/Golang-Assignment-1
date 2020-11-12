package main

import (
	//"fmt"
	"net/http"
	"html/template"
)

type pageData struct{
	Title string
	Name string
	Age int
	Phone int
	Company string
}

func myName(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("personal.html")
	data := pageData{Title: "Kloudone", Name : "Tamilarasan", Age: 24,Phone: 9787973192, Company : "Kloudone"}
	t.Execute(w,data)
	
}


func helloEveryone(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "tamil.html")
}


func main(){

	http.HandleFunc("/",  myName)
	http.HandleFunc("/tamil", helloEveryone)
	http.ListenAndServe(":9000",nil)


}