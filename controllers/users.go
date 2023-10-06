package controllers

import (
	"net/http"
)

type User struct{
	Template struct {
		New template
	}
}

func (u User) New(w http.ResponseWriter , r *http.Request){
	u.Template.New.Execute(w,nil)
}
