package model

import (
	"time"
	"github.com/go-sql-driver/mysql"
	"database/sql"
)

type Entry struct {
	Id int
	Author string
	Title string
	Body string
	DisplayAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   mysql.NullTime
}

type EntryModel struct {
	db sql.DB
}

func NewBlogModel(db sql.DB) *EntryModel {
	return &EntryModel{db}
}

func (bm *EntryModel) FindLatest(limit int) ([]*Entry, error) {
	rows, err := bm.db.Query(`
select
	id, title, author, body, display_at
from
	entry
where
	delete_at is NULL
order by
	display_at desc
limit
	?;
`, limit)
	if err != nil {
		return nil, err
	}

	es := []*Entry{}
	for rows.Next() {
		e := Entry{}
		err = rows.Scan(&e.Id, &e.Title, &e.Author, &e.Body, &e.DisplayAt)
		if err != nil {
			return nil, err
		}

		es = append(es, &e);
	}

	return es, nil
}