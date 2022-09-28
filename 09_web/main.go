package main

import ("fmt"
		"net/http")

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<p>Hello World</p>")
}

func contact(c http.ResponseWriter, r *http.Request){
	fmt.Fprintf(c,"<h1>contact us</h1>\n Now you can contact us on 08131388207")
}


func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	fmt.Println("starting server...")
	http.ListenAndServe(":3000",nil)
}
