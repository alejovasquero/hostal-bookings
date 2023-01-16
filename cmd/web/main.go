package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alejovasquero/hostal-bookings/config"
	"github.com/alejovasquero/hostal-bookings/http/routes"
	"github.com/alejovasquero/hostal-bookings/pkg/handler"
	"github.com/alejovasquero/hostal-bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var App config.AppConfig
var session *scs.SessionManager

func main() {
	App = config.AppConfig{}
	App.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	temp, err := render.CacheAllTemplates()

	if err != nil {
		log.Fatal(err)
	}

	App.TemplateCache = temp
	App.UseCache = false

	render.NewtTemplateRenderer(&App)

	repo := handler.NewHttpTemplateRepository(&App)
	handler.NewTemplateHandler(repo)

	//http.HandleFunc("/", repo.Index)
	//http.HandleFunc("/about", repo.About)
	//http.HandleFunc("/test", repo.Test)

	fmt.Println("Started server on port " + port)

	//err = http.ListenAndServe(port, nil)
	server := &http.Server{
		Addr:    port,
		Handler: routes.RoutesWithChi(&App),
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println("Error in the server startup")
		return
	}
}
