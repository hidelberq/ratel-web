package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type SoundCloudTrack struct {
	TrackId     int
	Name        string
	Title       string
	Author      string
	Description sql.NullString
	DisplayAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   mysql.NullTime
}

type SoundcloudTrackModel struct {
	db sql.DB
}

func NewSoundCloudTrackModel(db sql.DB) *SoundcloudTrackModel {
	return &SoundcloudTrackModel{db}
}

func (m *SoundcloudTrackModel) FindLatest(limit int) []*SoundCloudTrack {
	rows, err := m.db.Query(`
select
	track_id, name, title, author, description, display_at, created_at, updated_at
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

	ts := []*SoundCloudTrack{}
	for rows.Next() {
		t := SoundCloudTrack{}
		err = rows.Scan(&t.TrackId, &t.Name, &t.Title, &t.Author, &t.Description, &t.DisplayAt, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		ts = append(ts, &t)
	}

	return ts
}
