package model

import "time"

type Blog struct {
	ID        string
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
}

type BlogToSave struct {
	Title   string
	Content string
	Author  string
}
