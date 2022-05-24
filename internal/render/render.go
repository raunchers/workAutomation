package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// functions Map of functions usable in a template
var functions = template.FuncMap{}

// RenTemplate renders templates using html/template
func RenTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}

	// If template exist, t has a value and ok == true
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln(err)
	}

	// bytes buffer
	buf := new(bytes.Buffer)

	// take template, execute, don't pass any data and store in buf variable
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template,", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Map ; Template cache.
	myCache := map[string]*template.Template{}

	// Go to templates, get all files that have .page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//
	for _, page := range pages {
		// Extract name from cached templates (I.E home.page.tmpl)
		name := filepath.Base(page)
		// Template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Look for any file in the directory that ends in layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// add template set to the cache
		myCache[name] = ts
	}

	return myCache, nil
}
