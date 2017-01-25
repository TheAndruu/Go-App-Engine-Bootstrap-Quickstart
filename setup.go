package bootstrapquickstart

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates = make(map[string]*template.Template)

func init() {
	initializeTemplates()
	defineRoutes()
}

func defineRoutes() {
	http.HandleFunc("/showcase", showcaseHandler)
	http.HandleFunc("/", welcomeHandler)
}

// Base template is 'theme.html'  Can add any variety of content fillers in /layouts directory
func initializeTemplates() {
	layouts, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "templates/layouts/theme.html"))
	}
}

// A Welcome message with title, demonstrates passing data to a template
type Welcome struct {
	Title   string
	Message string
}

// A template taking a struct pointer (&message) containing data to render
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := Welcome{Title: "Bootstrap, Go, and GAE", Message: "Bootstrap added to Golang on App Engine.  Feel free to customize further"}

	// outerTheme refernces the template defined within theme.html
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}

func showcaseHandler(w http.ResponseWriter, r *http.Request) {
	// showcase.html doesn't expect any data, so pass empty string ""
	templates["showcase.html"].ExecuteTemplate(w, "outerTheme", "")
}
