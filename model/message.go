package model

import "database/sql"

type Message struct {
	Name string
	Email string
	Body string
}

type MessageModel struct {
	db sql.DB
}

func NewMessageModel(db sql.DB) *MessageModel {
	return &MessageModel{db}
}

func (mm *MessageModel) Create(m *Message) error {
	r, err := mm.db.Exec(`
		insert
		into
			message (name, email, body)
		values
			(?, ?, ?)`,
			m.Name,
			m.Email,
			m.Body,
	)
	if err != nil {
		return err
	}

	_, err = r.LastInsertId()
	return err
}