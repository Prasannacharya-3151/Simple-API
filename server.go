package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "handling the inoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handle users")
	})
	port := 3000
	fmt.Println("server is running on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil) //nil means user DefaultServeMux
}