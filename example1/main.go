package main

import (
	"fmt"
	"net/http"
)

func main() {

	cliexample()
}


func webServiceExample() {
	fmt.Println("Hello word")	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "web services are running")
	})	
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"./home.html")
	})
	http.ListenAndServe(":3000", nil)
}