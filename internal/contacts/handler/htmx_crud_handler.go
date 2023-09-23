package handler

import (
	"github.com/gin-gonic/gin"

	"htmx-go/internal/contacts/repository"
)

type ContactV2Handler struct {
	Repo repository.ContactRepository
}

func (h *ContactV2Handler) Index(c *gin.Context) {
	c.HTML(200, "index_v2.html", gin.H{})
}

func (h *ContactV2Handler) DisplayContacts(c *gin.Context) {
	keyword := c.DefaultQuery("q", "")
	contacts, _ := h.Repo.List(keyword)
	c.HTML(200, "contacts.html", gin.H{
		"Contacts": contacts,
	})
}
