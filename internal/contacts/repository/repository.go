package repository

import "htmx-go/internal/contacts"

type ContactRepository interface {
  List(keyword string) ([]contacts.Contact, error)
}

