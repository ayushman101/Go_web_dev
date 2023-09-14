package main

import (
	"fmt"
	"net/http"
)


func path1(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Welcome to the server</h1>")
}

func path2(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprint(w,"<h1>New Page</h1>")
}

func NotFoundPath(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>ERROR 404!</h1>Page Not Found")
}

type Router struct{}

func (router Router) ServeHTTP (w http.ResponseWriter, r *http.Request){
		switch r.URL.Path {
		case "/":
			path1(w,r)
		case "/newpage":
			path2(w,r)
		default: 
			NotFoundPath(w,r)
		}
	}
	
func main(){

	var router Router

	fmt.Println("The server is starting on port 3000")
	http.ListenAndServe(":3000",router)
	
}
