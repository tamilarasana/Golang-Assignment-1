package main

import ("fmt"
		"os"
		"bufio"
		"io/ioutil"
)



func main ()  {
	data, err := ioutil.ReadFile("new.txt")
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f,err := os.Open("new.txt")
	defer f.Close()

	if err != nil{
		fmt.Println(err)
	}
	reader  := bufio.NewReader(f)
	b1,err := reader.Peek(10)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b1))
}