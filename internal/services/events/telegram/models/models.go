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
	TextType  string       `db:"text_type"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type SystemDesignMsg struct {
	Title string
	Text  string
}

type Task struct {
	Id            int64        `db:"id"`
	TasksText     string       `db:"tasks_text"`
	Title         string       `db:"title"`
	TasksSolution string       `db:"tasks_solution"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}
