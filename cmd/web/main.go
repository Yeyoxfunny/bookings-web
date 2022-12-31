package main

import (
  "github.com/Yeyoxfunny/bookings-web/pkg/config"
  "github.com/Yeyoxfunny/bookings-web/pkg/handlers"
  "github.com/Yeyoxfunny/bookings-web/pkg/render"
  "github.com/alexedwards/scs/v2"
  "log"
  "net/http"
  "time"
)

var session *scs.SessionManager

func main() {
  tc, err := render.CreateTemplateCache()
  if err != nil {
    log.Fatalln("Cannot create template cache", err)
    return
  }

  var app config.AppConfig
  app.TemplateCache = tc
  app.UseCache = false
  app.InProduction = false

  session = scs.New()
  session.Lifetime = 24 * time.Hour
  session.Cookie.Persist = true
  session.Cookie.SameSite = http.SameSiteLaxMode
  session.Cookie.Secure = app.InProduction
  app.Session = session

  render.NewTemplates(&app)
  repository := handlers.NewRepo(&app)
  handlers.NewHandlers(repository)

  srv := &http.Server{
    Addr:    ":8080",
    Handler: routes(&app),
  }

  if err = srv.ListenAndServe(); err != nil {
    log.Fatalln(err)
  }
}
