package model

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Entry struct {
	Id        int
	Author    string
	Title     string
	Body      string
	DisplayAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt mysql.NullTime
}

type EntryModel struct {
	db sql.DB
}

func NewEntryModel(db sql.DB) *EntryModel {
	return &EntryModel{db}
}

func (em *EntryModel) FindLatest(limit int) ([]*Entry, error) {
	rows, err := em.db.Query(`
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

		es = append(es, &e)
	}

	return es, nil
}

func (em *EntryModel) FindAll(limit int) ([]*Entry, error) {
	stmt, err := em.db.Prepare(`
select
	id, title, author, body, display_at, delete_at
from
	entry
order by
	display_at desc
limit
	?;`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}

	es := []*Entry{}
	for rows.Next() {
		e := Entry{}
		err = rows.Scan(&e.Id, &e.Title, &e.Author, &e.Body, &e.DisplayAt, &e.DeletedAt)
		if err != nil {
			return nil, err
		}

		es = append(es, &e)
	}

	return es, nil
}

func (em *EntryModel) FindById(id int) (*Entry, error) {
	stmt, err := em.db.Prepare(`
select
	id, title, author, body, display_at, delete_at
from
	entry
where
	id = ?`)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	e := &Entry{}
	err = row.Scan(&e.Id, &e.Title, &e.Author, &e.Body, &e.DisplayAt, &e.DeletedAt)
	return e, err
}

func (em *EntryModel) Create(e *Entry) (sql.Result, error) {
	stmt, err := em.db.Prepare(`
insert into
	entry
	(author, title, body, display_at, delete_at)
values
	(?, ?, ?, ?, ?);
	`)
	if err != nil {
		return nil, err
	}

	ret, err := stmt.Exec(e.Title, e.Author, e.Body, e.DisplayAt, e.DeletedAt)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (em *EntryModel) Update(e *Entry, id int) (sql.Result, error) {
	stmt, err := em.db.Prepare(`
update
	entry
set
	author = ?,
	title = ?,
	body = ?,
	display_at = ?,
	delete_at = ?
where
	id = ?;`)
	if err != nil {
		return nil, err
	}

	return stmt.Exec(
		e.Author,
		e.Title,
		e.Body,
		e.DisplayAt,
		e.DeletedAt,
		id)
}
