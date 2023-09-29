package controllers

import (
	"net/http"
	"githum.com/ayushman101/Go_web_dev/views"
	"github.com/go-chi/chi/v5"
)


func StaticHandler(t views.Template, data interface{}) http.HandlerFunc{
	
	return func(w http.ResponseWriter, r * http.Request){
		 t.Execute(w,data)
	}
}


func Userpage(t views.Template, w http.ResponseWriter, r *http.Request){


	userID:=chi.URLParam(r,"userID")

	ctx:= r.Context()
	key:= ctx.Value("userID")

	user:= struct{
		
		Name string
		Userid any
		Key any
		Test []string
		
	}{
		
		Name: "John Smith",
		Userid : userID,
		Key: key,
		Test: []string{
			"red",
			"yellow",
		},
	}

	t.Execute(w,user)
	
}
