package handlers

import (
	"net/http"

	"github.com/singhudeek/booking/pkg/config"
	"github.com/singhudeek/booking/pkg/models"
	"github.com/singhudeek/booking/pkg/renders"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repos
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository or the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the abount page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again."
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// render the Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// render the fenerals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// render the majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// render the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// render the Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
