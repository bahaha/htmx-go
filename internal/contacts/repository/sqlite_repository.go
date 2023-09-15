package repository

import (
	"database/sql"
	"htmx-go/internal/contacts"
)

type SqliteContactRepository struct {
  DB *sql.DB
}

func (r *SqliteContactRepository) List() ([]contacts.Contact, error) {
  rows, err := r.DB.Query("SELECT id, first_name, last_name, phone, email FROM contacts")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var res []contacts.Contact
  for rows.Next() {
    var contact contacts.Contact
    err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
    if err != nil {
      return nil, err
    }
    res = append(res, contact)
  }

  return res, nil
}
