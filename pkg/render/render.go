package render

import (
  "bytes"
  "github.com/Yeyoxfunny/bookings-web/pkg/config"
  "github.com/Yeyoxfunny/bookings-web/pkg/models"
  "html/template"
  "log"
  "net/http"
  "path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
  app = a
}

func addDefaultData(data *models.TemplateData) *models.TemplateData {
  return data
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data *models.TemplateData) {
  var tCache map[string]*template.Template
  if app.UseCache {
    tCache = app.TemplateCache
  } else {
    tCache, _ = CreateTemplateCache()
  }

  tmpl, exists := tCache[tmplName]
  if !exists {
    log.Fatalf("Template with name %v does not exits\n", tmplName)
    return
  }

  buf := new(bytes.Buffer)

  data = addDefaultData(data)
  _ = tmpl.Execute(buf, data)

  _, err := buf.WriteTo(w)
  if err != nil {
    log.Fatalln("Error writing template to browser", err)
  }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
  pages, err := filepath.Glob("./templates/*.page.gohtml")
  if err != nil {
    return nil, err
  }

  tCache := make(map[string]*template.Template)
  for _, page := range pages {
    name := filepath.Base(page)
    tmpl, err := template.New(name).ParseFiles(page)
    if err != nil {
      return nil, err
    }
    matches, err := filepath.Glob("./templates/*.layout.gohtml")
    if err != nil {
      return nil, err
    }

    if len(matches) > 0 {
      tmpl, err = tmpl.ParseGlob("./templates/*.layout.gohtml")
      if err != nil {
        return nil, err
      }
    }

    tCache[name] = tmpl
  }
  return tCache, nil
}
