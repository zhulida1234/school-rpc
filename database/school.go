package database

import "gorm.io/gorm"

type SchoolDB struct {
	gorm *gorm.DB
}

func NewSchoolDB(gorm *gorm.DB) *SchoolDB {
	return &SchoolDB{
		gorm: gorm,
	}
}
