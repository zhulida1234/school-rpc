package database

import "gorm.io/gorm"

type SchoolDB struct {
	gorm *gorm.DB
}

func (schoolDB *SchoolDB) Close() error {
	sql, err := schoolDB.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func NewSchoolDB(gorm *gorm.DB) *SchoolDB {
	return &SchoolDB{
		gorm: gorm,
	}
}
