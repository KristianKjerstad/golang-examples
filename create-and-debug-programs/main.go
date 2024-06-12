package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func main() {
	// simpleCLI()

	http.HandleFunc("/", webServiceHandler)
	http.ListenAndServe(":3000", nil)
}



func simpleCLI() {
	fmt.Println("What should I scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	fmt.Println(s + "!")
}


func webServiceHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)


}