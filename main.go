package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
)
func logHttpHeader(r *http.Request){

	fmt.Println("Method: ",r.Method," | Proto: ", r.Proto)
	fmt.Println("Host : ", r.Host)
}

func TemplateExecute(w http.ResponseWriter, filepath string, data any){

	tpl,err:=template.ParseFiles(filepath)
	if err!=nil{
	
		http.Error(w,"Failed to parse response html", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = tpl.Execute(w,data)

	if err!=nil{

		http.Error(w,"Failed to parse response html", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}


}

func path1(w http.ResponseWriter, r *http.Request){
	logHttpHeader(r)	
	TemplateExecute(w,"./templates/home.gohtml",nil)
	//fmt.Fprint(w,"<h1>Welcome to the server</h1>")
}

func path2(w http.ResponseWriter, r *http.Request){
	logHttpHeader(r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	
	TemplateExecute(w,"./templates/newpage.gohtml",nil)

	//fmt.Fprint(w,"<h1>New Page 2</h1>")
}


func userPath(w http.ResponseWriter, r *http.Request){
	logHttpHeader(r)
	userID:=chi.URLParam(r,"userID")

	ctx := r.Context()
  	key := ctx.Value("userID")

  // respond to the client
  //w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
	
  user:=struct{
	Name string
	Userid any
	Key any
	Test []string
  }{
	Name: "John Smith",
	Userid: userID,
	Key: key,
	Test: []string{
		"red",
		"yellow",
	},
  }

  	TemplateExecute(w,"./templates/user.gohtml",user)
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
