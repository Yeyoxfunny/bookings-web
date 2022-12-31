package handlers

import (
  "github.com/Yeyoxfunny/bookings-web/pkg/config"
  "github.com/Yeyoxfunny/bookings-web/pkg/models"
  "github.com/Yeyoxfunny/bookings-web/pkg/render"
  "log"
  "net/http"
)

var Repo *Repository

type Repository struct {
  App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
  return &Repository{
    App: a,
  }
}

func NewHandlers(r *Repository) {
  Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
  remoteIp := r.RemoteAddr
  log.Println("Remote ip address", remoteIp)
  m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
  render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
  stringMap := make(map[string]string)
  stringMap["test"] = "Hello, again"
  stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

  render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
    StringMap: stringMap,
  })
}
