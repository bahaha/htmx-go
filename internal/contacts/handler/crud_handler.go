package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"htmx-go/internal/contacts"
	"htmx-go/internal/contacts/repository"
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
		"Keyword":  keyword,
		"Contacts": contacts,
	})
}

func (h *ContactHandler) findContact(pathId string) (*contacts.Contact, error) {
	id, err := strconv.Atoi(pathId)
	if err != nil {
		return nil, err
	}

	return h.Repo.Find(id)
}

func (h *ContactHandler) FindContact(c *gin.Context) {
	contact, err := h.findContact(c.Param("id"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.HTML(200, "show.html", gin.H{
		"Contact": contact,
	})
}

func (h *ContactHandler) ContactViewToEdit(c *gin.Context) {
	contact, err := h.findContact(c.Param("id"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.HTML(200, "edit.html", gin.H{
		"Contact": contact,
	})
}

func (h *ContactHandler) EditContact(c *gin.Context) {
	var contact contacts.Contact
	if err := c.ShouldBind(&contact); err != nil {
		// TODO: bad request with htmx
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	contact.ID, _ = strconv.Atoi(id)

	err := h.Repo.Update(&contact)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/contacts/%d", contact.ID))
}
