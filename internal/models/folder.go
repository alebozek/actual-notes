package models

import "time"

type Folder struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FolderNote struct {
	FolderID uint
	NoteID   uint
}
