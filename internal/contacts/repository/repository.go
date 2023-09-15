package repository

import "htmx-go/internal/contacts"

type ContactRepository interface {
  List() ([]contacts.Contact, error)
}
