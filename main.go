package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Your code to handle the HTTP request goes here
	var a string = register()
	fmt.Fprintln(a)
}
func main() {
	http.HandleFunc("/", handler) // Register the handler function for the root URL

	port := 8080 // Choose the port you want to listen on
	fmt.Printf("Server is listening on port %d...\n", port)
	resp,err = http://20.244.56.144/train/register
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // Start the HTTP server
	if err != nil {
		fmt.Println("Error:", err)
	}


}
func register(){
	companyname := "Train central"
	ownerName:= "Sujan"
	rollNo:=122010331037
	ownerEmail:"saisujan2002@gmail.com"
	accessCode: "FKDLjg"
}

func response(){
	companyName:"Train Central"
	var clientID string
	var clietSecret
	fmt.Println(companyName)
	fmt.println(clientID)
	fmt.println(clietSecret)

})