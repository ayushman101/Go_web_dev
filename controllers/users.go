package controllers

import (

	"net/http"
	"githum.com/ayushman101/Go_web_dev/views"
)

type User struct{
	Template struct {
		New views.Template
	}
}

func (u User) New(w http.ResponseWriter , r *http.Request){
	u.Template.New.Execute(w,nil)
}
