package handler

import (
	"net/http"

	"github.com/alejovasquero/hostal-bookings/config"
	"github.com/alejovasquero/hostal-bookings/models"
	"github.com/alejovasquero/hostal-bookings/pkg/render"
)

var Repo *HttpTemplateRepository

type HttpTemplateRepository struct {
	App *config.AppConfig
}

func NewHttpTemplateRepository(a *config.AppConfig) *HttpTemplateRepository {
	return &HttpTemplateRepository{
		App: a,
	}
}

func NewTemplateHandler(r *HttpTemplateRepository) {
	Repo = r
}

func (hr *HttpTemplateRepository) Index(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	hr.App.Session.Put(r.Context(), "ip-address", remoteIp)
	render.WriteTemplateFromFullCache("complete.page.html", w, &models.TemplateData{})
}

func (hr *HttpTemplateRepository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, my friend!"

	ipAddress := hr.App.Session.GetString(r.Context(), "ip-address")
	stringMap["remote-ip"] = ipAddress

	render.WriteTemplateFromFullCache("about.page.html", w, &models.TemplateData{
		StringMap: stringMap,
	})
}

func (hr *HttpTemplateRepository) Test(w http.ResponseWriter, r *http.Request) {
	render.WriteTemplateFromFullCache("complete.page.html", w, &models.TemplateData{})
}
