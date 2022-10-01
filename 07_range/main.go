package main

import "fmt"

func main(){
	grades := []int{10,20,30,40,50}
	for i,id := range grades{
		fmt.Printf("index %d value = %d\n",i,id)
	}
	students := map[int] string{1:"Biodun Akinlawon",2:"Gbemisola Afolabi"}
	for _,v := range students{
		fmt.Printf(" value = %s\n",v)
	}

}