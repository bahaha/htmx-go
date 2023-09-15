package main

import (
	"htmx-go/internal/contacts/handler"
	"htmx-go/internal/contacts/repository"
	"htmx-go/internal/database"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main () {
  r := gin.Default()
  r.HTMLRender = loadTemplates("web/templates")
  handlers := initHandlers()
  r.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/contacts")
  })
  r.GET("/contacts", handlers.ListContacts)
  r.GET("/contacts/:id", handlers.FindContact)

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

func loadTemplates(templateDir string) multitemplate.Renderer {
  renderer := multitemplate.NewRenderer()

  layouts, err := filepath.Glob(templateDir + "/layout/*.html")
  if err != nil {
    panic(err.Error())
  }

  for _, layout := range layouts {
    layoutFileName := filepath.Base(layout)
    layoutName := layoutFileName[:len(layoutFileName)-len(filepath.Ext(layoutFileName))]

    pages, err := filepath.Glob(templateDir + "/pages/" + layoutName + "/*.html")
    if err != nil {
      panic(err.Error())
    }

    for _, page := range pages {
      pageName := filepath.Base(page)
      renderer.AddFromFiles(pageName, layout, page)
    }
  }

  return renderer
}
