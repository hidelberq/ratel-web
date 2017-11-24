package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Track struct {
	TrackId     int
	Name        string
	Title       string
	Author      string
	Description string
	DisplayAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   mysql.NullTime
}

type TrackModel struct {
	db sql.DB
}

func NewTrackModel(db sql.DB) *TrackModel {
	return &TrackModel{db}
}

func (m *TrackModel) FindLatest(limit int) []*Track {
	rows, err := m.db.Query(`
select
	track_id, name, title, author, description, display_at, created_at, updated_at, deleted_at
from
	soundcloud_track
where
	deleted_at is NULL
order by
	display_at desc
limit
	?;`, limit)
	if err != nil {
		panic(err.Error())
	}

	ts := []*Track{}
	for rows.Next() {
		t := Track{}
		err = rows.Scan(&t.TrackId, &t.Name, &t.Title, &t.Author, &t.Description, &t.DisplayAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			panic(err.Error())
		}

		ts = append(ts, &t)
	}

	return ts
}

func (m *TrackModel) FindById(id int) (*Track, error) {
	row := m.db.QueryRow(`
select
	track_id, name, title, author, description, display_at, created_at, updated_at, deleted_at
from
	soundcloud_track
where
	track_id = ?;`, id)
	if row == nil {
		return nil, nil
	}

	t := &Track{}
	if err := row.Scan(&t.TrackId, &t.Name, &t.Title, &t.Author, &t.Description, &t.DisplayAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt); err != nil {
		return nil, err
	}

	return t, nil

}
