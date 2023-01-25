package handler

import (
	"encoding/json"
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
	render.WriteTemplateFromFullCache("index.page.html", r, w, &models.TemplateData{})
}

func (hr *HttpTemplateRepository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, my friend!"

	ipAddress := hr.App.Session.GetString(r.Context(), "ip-address")
	stringMap["remote-ip"] = ipAddress

	render.WriteTemplateFromFullCache("about.page.html", r, w, &models.TemplateData{
		StringMap: stringMap,
	})
}

func (hr *HttpTemplateRepository) Contact(w http.ResponseWriter, r *http.Request) {
	render.WriteTemplateFromFullCache("contact.page.html", r, w, &models.TemplateData{})
}

func (hr *HttpTemplateRepository) Jigsaw(w http.ResponseWriter, r *http.Request) {
	render.WriteTemplateFromFullCache("jigsaw.page.html", r, w, &models.TemplateData{})
}

func (hr *HttpTemplateRepository) Torture(w http.ResponseWriter, r *http.Request) {
	render.WriteTemplateFromFullCache("torture.page.html", r, w, &models.TemplateData{})
}

func (hr *HttpTemplateRepository) SearchAvalability(w http.ResponseWriter, r *http.Request) {
	render.WriteTemplateFromFullCache("search-availability.page.html", r, w, &models.TemplateData{})
}

type jsonResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (hr *HttpTemplateRepository) PostAvalability(w http.ResponseWriter, r *http.Request) {
	startDate := r.FormValue("start")
	endDate := r.FormValue("end")

	//w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", startDate, endDate)))

	response := &jsonResponse{
		Status:    200,
		Message:   "Request Successfull",
		StartDate: startDate,
		EndDate:   endDate,
	}
	marshal, err := json.MarshalIndent(response, "", "    ")

	if err == nil {
		w.Write(marshal)
		w.Header().Add("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
