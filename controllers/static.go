package controllers

import (
	"net/http"
	"githum.com/ayushman101/Go_web_dev/views"
)


func StaticHandler(t views.Template) http.HandlerFunc{
	
	return func(w http.ResponseWriter, r * http.Request){
		 t.Execute(w,nil)
	}
}
