package main

import (
  "fmt"
  "regexp"
)

func main() {
  re := regexp.MustCompile("et$")
  
  fmt.Println(re.FindString("cricket"))
  fmt.Println(re.FindString("hacked"))
  fmt.Println(re.FindString("racket"))

  //Find the  index  method
  
  rl := regexp.MustCompile("tel")
  fmt.Println(rl.FindStringIndex("telephone"))
  fmt.Println(rl.FindStringIndex("carpet"))
  fmt.Println(rl.FindStringIndex("cartel"))
}