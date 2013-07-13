/**
 * User: Jackong
 * Date: 13-6-29
 * Time: 下午10:40
 */
package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"errors"
)


type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	fileName := dir + p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := dir + title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

const (
	lenPath = len("/view/")
	dir = "golearn/"
)

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	render(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	render(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r * http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func render(w http.ResponseWriter, tpl string, p *Page) {
	err := tpls.ExecuteTemplate(w, tpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var (
tpls = template.Must(template.ParseFiles(dir + "edit.html", dir + "view.html"))
titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

func getTitle(w  http.ResponseWriter, r *http.Request) (title string, err error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		http.NotFound(w, r)
		err = errors.New("Invalid page title")
	}
	return
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe("", nil)
}
