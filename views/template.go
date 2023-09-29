package views

import (
	"log"
	"html/template"
	"net/http"
	"fmt"
	"io/fs"
)



func ParseFS (filesystem fs.FS, patterns ...string) (Template, error) {

	tpl, err:= template.ParseFS(filesystem, patterns...)

	if err!=nil {
		return Template{}, fmt.Errorf("Parsing Template error %w ", err)
	}

	return Template{
		htmlTpl: tpl,
	},nil
}


func Parse (filepath string) (Template,error){

	t,err:= template.ParseFiles(filepath)

	if err!=nil{

		return Template{},fmt.Errorf("Parsing Template %w",err)
	}

	return Template{
		htmlTpl: t,
	},nil
}

type Template struct{

	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	err:= t.htmlTpl.Execute(w,data)

	if err!=nil{
		
		log.Printf("Executing template: %v", err)
		http.Error(w,"There was an error executing the template", http.StatusInternalServerError)
		return
	}
}  
