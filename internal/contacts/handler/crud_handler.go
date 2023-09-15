package handler

import (
  "github.com/gin-gonic/gin"
  // "html/template"
  // "htmx-go/internal/contacts"
)

func ListContacts(c *gin.Context) {
  c.HTML(200, "index.tmpl", gin.H{})
}
