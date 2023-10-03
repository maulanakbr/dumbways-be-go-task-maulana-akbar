package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name   string
	Vision string
}
