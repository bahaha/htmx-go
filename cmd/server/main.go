package main

import (
	"html/template"
	"htmx-go/internal/contacts/handler"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main () {
  r := gin.Default()
  r.SetHTMLTemplate(loadTemplates("web/templates"))
  r.GET("/", handler.ListContacts)

  r.Run(":55688")
}

func loadTemplates(templateDir string) *template.Template {
  files, err := filepath.Glob(filepath.Join(templateDir, "*.tmpl"))
  if err != nil {
    panic(err)
  }

  return template.Must(template.New("").ParseFiles(files...))
}
