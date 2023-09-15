package repository

import (
	"database/sql"
	"htmx-go/internal/contacts"
)

type SqliteContactRepository struct {
	DB *sql.DB
}

func (r *SqliteContactRepository) List(keyword string) ([]contacts.Contact, error) {
	var rows *sql.Rows
	var err error

	if keyword != "" {
		keywordPattern := "%" + keyword + "%"
		rows, err = r.DB.Query(
			"SELECT id, first_name, last_name, phone, email FROM contacts WHERE first_name LIKE ? OR last_name LIKE ? OR phone LIKE ? OR email LIKE ?",
			keywordPattern, keywordPattern, keywordPattern, keywordPattern,
		)
	} else {
		rows, err = r.DB.Query("SELECT id, first_name, last_name, phone, email FROM contacts")
	}

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

func (r *SqliteContactRepository) Find(id int) (*contacts.Contact, error) {
  row := r.DB.QueryRow("SELECT id, first_name, last_name, phone, email FROM contacts WHERE id = ?", id)

  var contact contacts.Contact
  err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
  if err != nil {
    return nil, err
  }

  return &contact, nil
}

