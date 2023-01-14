package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/alejovasquero/hostal-bookings/config"
	"github.com/alejovasquero/hostal-bookings/models"
)

var appConfig *config.AppConfig

func NewtTemplateRenderer(conf *config.AppConfig) {
	appConfig = conf
}

func WriteTemplate(templateName string, w http.ResponseWriter) {
	loadTemplate, err := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.html")
	if err != nil {
		fmt.Println("Error in the tamplate load: ", err)
		return
	}

	err = loadTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error loading template ", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func WriteTemplateCache(templateName string, w http.ResponseWriter) {
	var loadTemplate *template.Template
	var err error

	_, isPresent := templateCache[templateName]

	if !isPresent {
		err = readTemplateToCache(templateName)
		if err != nil {
			fmt.Println("Error reading template: ", err)
			return
		}
	}

	loadTemplate = templateCache[templateName]

	err = loadTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error loading template ", err)
		return
	}
}

func WriteTemplateFromFullCache(templateName string, w http.ResponseWriter, data *models.TemplateData) {
	var templates map[string]*template.Template
	var err error

	if !appConfig.UseCache {
		templates, err = CacheAllTemplates()

		if err != nil {
			log.Println("Template not found:", templateName)
			return
		}
	} else {
		templates = appConfig.TemplateCache
	}

	var ok bool
	finalTemplate, ok := templates[templateName]

	if !ok {
		log.Println("Template not found:", templateName)
		return
	}

	buffer := new(bytes.Buffer)

	data = addDefaultTemplateData(data)

	err = finalTemplate.Execute(buffer, data)

	if err != nil {
		log.Println(err)
	}

	numberOfData, err := buffer.WriteTo(w)
	log.Printf("Wrote %d bytes to response writer: template %s\n", numberOfData, templateName)

	if err != nil {
		log.Println(err)
	}
}

func readTemplateToCache(templateName string) error {
	fmt.Printf("Reading template from disk: %s\n", templateName)

	loadTemplate, err := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.html")
	if err != nil {
		fmt.Println("Error in the template load: ", err)
		return err
	}
	templateCache[templateName] = loadTemplate
	return nil
}

func addDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

func CacheAllTemplates() (map[string]*template.Template, error) {
	loadedTemplates := map[string]*template.Template{}

	files, err := filepath.Glob("./templates/*.html")

	if err != nil {
		return loadedTemplates, err
	}

	for _, file := range files {
		fmt.Println("Loaded template: ", file)
		fileName := filepath.Base(file)

		loadTemplate, err := template.ParseFiles(file)

		if err != nil {
			return loadedTemplates, err
		}

		loadedTemplates[fileName] = loadTemplate

		filenames, err := filepath.Glob("./templates/*.layout.html")

		if err == nil && len(filenames) > 0 {
			loadTemplate.ParseGlob("./templates/*.layout.html")
		} else {
			log.Println("No layouts found")
		}
	}

	return loadedTemplates, nil
}
