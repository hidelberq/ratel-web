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
	DisplayTime time.Time
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

func (m *SoundcloudTrackModel) FindAll() []*SoundCloudTrack {
	rows, err := m.db.Query(`
	select track_id, name, title, author, description, display_time, created_at, updated_at
	from soundcloud_track
	where deleted_at is NULL;`)
	if err != nil {
		panic(err.Error())
	}

	ts := []*SoundCloudTrack{}
	for rows.Next() {
		t := SoundCloudTrack{}
		err = rows.Scan(&t.TrackId, &t.Name, &t.Title, &t.Author, &t.Description, &t.DisplayTime, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		ts = append(ts, &t)
	}

	return ts
}
