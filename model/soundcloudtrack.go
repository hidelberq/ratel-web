package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
	"log"
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
	deleted_at is null
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

func (m *TrackModel) FindAll(limit int) ([]*Track, error) {
	stmt, err := m.db.Prepare(
		`
select
	track_id, name, title, author, description, display_at, created_at, updated_at, deleted_at
from
	soundcloud_track
order by
	display_at desc
limit
	?;
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}

	ts := []*Track{}
	for rows.Next() {
		t := Track{}
		err = rows.Scan(&t.TrackId, &t.Name, &t.Title, &t.Author, &t.Description, &t.DisplayAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, err
		}

		ts = append(ts, &t)
	}

	return ts, nil
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

func (m *TrackModel) Update(id int, track *Track) error {
	_, err := m.db.Exec(`
update
	soundcloud_track
set
	track_id = ?,
	name = ?,
	title = ?,
	author = ?,
	description = ?,
	display_at = ?,
	updated_at = ?,
	deleted_at = ?
where
	track_id = ?;`,
		track.TrackId,
		track.Name,
		track.Title,
		track.Author,
		track.Description,
		track.DisplayAt,
		track.UpdatedAt,
		track.DeletedAt,
		id)
	return err
}

func (m *TrackModel) Create(track *Track) error {
	ret, err := m.db.Exec(`
insert into
	soundcloud_track
	(
	track_id,
	name,
	title,
	author,
	description,
	display_at,
	updated_at
	)
values
	(?, ?, ?, ?, ?, ?, ?);
`,
		track.TrackId,
		track.Name,
		track.Title,
		track.Author,
		track.Description,
		track.DisplayAt,
		track.UpdatedAt)
	log.Println(ret, err)
	return err
}
