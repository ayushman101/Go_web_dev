package main

import (
	"fmt"
	"net/http"
)

func main(){

	handlerfunc:=func(w http.ResponseWriter, _ *http.Request){
		fmt.Fprint(w,"<h1>Welcome to the Server</h1>")
	}
	
	http.HandleFunc("/",handlerfunc)

	fmt.Println("The server is starting on port 3000")
	http.ListenAndServe(":3000",nil)
	
}
