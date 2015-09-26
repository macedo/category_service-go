package model

import "time"

type Category struct {
	ID        int
	ParentID  int
	Lft       int
	Rgt       int
	Depth     int
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Categories []Category
