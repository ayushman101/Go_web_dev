package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func path1(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Method: ",r.Method," | Proto: ", r.Proto)
	fmt.Println("Host : ", r.Host)
	fmt.Fprint(w,"<h1>Welcome to the server</h1>")
}

func path2(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprint(w,"<h1>New Page 2</h1>")
}


func userPath(w http.ResponseWriter, r *http.Request){

	userID:=chi.URLParam(r,"userID")

	ctx := r.Context()
  	key := ctx.Value("userID")

  // respond to the client
  w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
	
  	//fmt.Fprint(w,"<h1>User: ", userID,"</h1> <h2> Key: ",key," </h2>")
	
}

func NotFoundPath(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>ERROR 404!</h1>Page Not Found")
}

func main(){

	r:=chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", path1)
	r.Get("/new-page",path2)
	r.Get("/users/{userID}",userPath)
	r.NotFound(NotFoundPath)

	fmt.Println("The server is starting on port 3000")
	http.ListenAndServe(":3000",r)
	
}
