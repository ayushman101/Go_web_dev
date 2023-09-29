package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
	"log"
	"githum.com/ayushman101/Go_web_dev/views"
	"githum.com/ayushman101/Go_web_dev/controllers"
	"githum.com/ayushman101/Go_web_dev/templates"
)

func logHttpHeader(r *http.Request){

	fmt.Println("Method: ",r.Method," | Proto: ", r.Proto)
	fmt.Println("Host : ", r.Host)
}

func TemplateExecute(w http.ResponseWriter, filepath string, data any){
	
	tpl,err:=views.Parse(filepath)
	if err!=nil{
	
		http.Error(w,"Failed to parse response html", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	tpl.Execute(w,data)

}


func userPath(w http.ResponseWriter, r *http.Request){

	tpl,err:= views.ParseFS(templates.FS,"layout.gohtml","user.gohtml")
	if err!=nil{
		panic(err)
	}

	controllers.Userpage(tpl,w,r)
	
}

func NotFoundPath(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>ERROR 404!</h1>Page Not Found")
}

func main(){

	r:=chi.NewRouter()
	
	tpl,err:= views.ParseFS( templates.FS, "layout.gohtml","home.gohtml")
	if err!=nil{
		panic(err)
	}

	r.Use(middleware.Logger)
	r.Get("/",controllers.StaticHandler(tpl,nil))
	
	tpl,err = views.ParseFS(templates.FS, "layout.gohtml","newpage.gohtml")
	if err!=nil{
		panic(err)
	}



	r.Get("/new-page",controllers.StaticHandler(tpl,nil))

	r.Get("/users/{userID}",userPath)
	
	r.NotFound(NotFoundPath)

	fmt.Println("The server is starting on port 3000")
	
	http.ListenAndServe(":3000",r)
	
}
