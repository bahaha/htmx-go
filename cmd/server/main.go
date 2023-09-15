package main

import (
	"html/template"
	"htmx-go/internal/contacts/handler"
	"htmx-go/internal/contacts/repository"
	"htmx-go/internal/database"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main () {
  r := gin.Default()
  r.SetHTMLTemplate(loadTemplates("web/templates"))
  handlers := initHandlers()
  r.GET("/", handlers.ListContacts)

  r.Run(":55688")
}

func initHandlers() *handler.ContactHandler {
  db, err := database.DataSource()
  if err != nil {
    panic(err)
  }

  repo := &repository.SqliteContactRepository{DB: db}
  return &handler.ContactHandler{Repo: repo}
}

func loadTemplates(templateDir string) *template.Template {
  files, err := filepath.Glob(filepath.Join(templateDir, "*.tmpl"))
  if err != nil {
    panic(err)
  }

  return template.Must(template.New("").ParseFiles(files...))
}
