package models

import (
	"database/sql"
	"time"
)

type GoInfoData struct {
	Data string `json:"data"`
}

type GoInfo struct {
	Id        int64        `db:"id"`
	Text      string       `db:"text"`
	Title     string       `db:"title"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
