package repository

import "htmx-go/internal/contacts"

type ContactRepository interface {
	List(keyword string) ([]contacts.Contact, error)
	Find(id int) (*contacts.Contact, error)
	Update(contact *contacts.Contact) error
	Delete(id int) error
}
