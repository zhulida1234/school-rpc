package database

type Clazz struct {
	Id    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"name"`
	No    string `json:"no"`
	Grade uint32 `json:"grade"`
}

func (schoolDB *SchoolDB) CreateClazz(clazz *Clazz) error {
	err := schoolDB.gorm.Create(clazz).Error
	if err != nil {
		return err
	}
	return nil
}
