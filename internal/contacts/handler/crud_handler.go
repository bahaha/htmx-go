package handler

import (
	"htmx-go/internal/contacts/repository"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
  Repo repository.ContactRepository
}

func (h *ContactHandler) ListContacts(c *gin.Context) {
  keyword := c.DefaultQuery("q", "")
  contacts, err := h.Repo.List(keyword)
  if err != nil {
    c.AbortWithError(500, err)
    return
  }

  c.HTML(200, "index.tmpl", gin.H{
    "Keyword": keyword,
    "Contacts": contacts,
  })
}
