package main

import "fmt"

func main(){
	emails := make(map[int]string)
	emails[1] = "biodunakinlawon@gmail.com"
	emails[2] = "gbemisolaafolabi@gmail.com"

	//another declaration method
	students := map[int] string{1:"Biodun Akinlawon",2:"Gbemisola Afolabi"}
	fmt.Println(emails[2]+" "+students[1])
}
