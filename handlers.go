package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title string
}

type BlogPage struct {
	Title     string
	PostStyle template.HTML
	Body      template.HTML
}

func RenderPost(w http.ResponseWriter, r *http.Request) {

	page := `<!DOCTYPE html>
          	<head>
          		<title>{{.Title}}</title>
              <link href="https://fonts.googleapis.com/css?family=Poppins:300,400,500,600,700" rel="stylesheet">
              {{.PostStyle}}
          	</head>
          	<body>
              <div class='container'>
                {{.Body}}
              </div>
            </body>
          </html>`

	input, err := ioutil.ReadFile("posts/truthless.md")
	if err != nil {
		log.Println("error opening post off of disk.")
	}

	output := blackfriday.MarkdownCommon(input)

	//s := string(byteArray[:n])

	//n := bytes.Index(byteArray, []byte{0})

	//html := bluemonday.UGCPolicy().SanitizeBytes(output)

	n := len(output)
	partialContent := string(output[:n])

	pageDefinition := new(BlogPage)
	//page = &Page{Title: "this is a test"}
	pageDefinition.Title = "testing !"
	pageDefinition.Body = template.HTML(partialContent)
	pageDefinition.PostStyle = template.HTML(BlogCSS)

	tmpl := template.New("page")

	if tmpl, err = tmpl.Parse(page); err != nil {
		fmt.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pageDefinition)
	if err != nil {
		fmt.Println(err)
	}

	result := buf.String()

	//fmt.Println(s)

	/*

		doc, err := goquery.NewDocument(result)

		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find("#metadata").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			//band := s.Find("a").Text()
			s.Text()
			//title := s.Find("i").Text()
			//fmt.Printf("Review %d: %ss\n", i, band)
		}) */

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))

	if _, err := w.Write([]byte(result)); err != nil {
		log.Println("unable to write html.")
	}

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
