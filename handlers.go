package main

import (
	"html/template"
	"io"
	"net/http"
)

type Page struct {
	Title string
}

// health check...
func Health(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

/*
func loadIndexPage() (*Page, error) {
	filename := "index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

func Root(w http.ResponseWriter, r *http.Request) {
	//p, err := loadIndexPage()
	//if err != nil {
	//p = &Page{Title: title}
	//}
	//var page Page
	p := new(Page)
	//page = &Page{Title: "this is a test"}
	p.Title = "testing !"
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

/*
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root route")

	//t, _ := template.ParseFiles(filename)

	//var result IndexData

	//result.Actions = make([]Action, len(actionData.Res))

	//t.Execute(w, result)

}*/
