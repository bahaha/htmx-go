package handler

import (
	"htmx-go/internal/contacts/repository"
	"strconv"

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

  c.HTML(200, "index.html", gin.H{
    "Keyword": keyword,
    "Contacts": contacts,
  })
}

func (h *ContactHandler) FindContact(c *gin.Context) {
  idStr := c.Param("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    c.AbortWithError(500, err)
    return
  }

  contact, err := h.Repo.Find(id)
  if err != nil {
    c.AbortWithError(500, err)
    return
  }

  c.HTML(200, "show.html", gin.H{
    "Contact": contact,
  })
}
