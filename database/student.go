package database

type Student struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name"`
	Age       uint32 `json:"age"`
	Gender    uint32 `json:"gender"`
	Mobile    string `json:"mobile"`
	ClassName string `json:"className"`
	Grade     uint32 `json:"grade"`
}

func (schoolDB *SchoolDB) FindStudentList(pageSize uint32, pageNo uint32) ([]Student, error) {
	var students []Student

	// 计算偏移量
	offset := (pageNo - 1) * pageSize

	// 使用 GORM 查询学生列表
	if err := schoolDB.gorm.Offset(int(offset)).Limit(int(pageSize)).Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (schoolDB *SchoolDB) CreateStudent(student *Student) error {
	err := schoolDB.gorm.Create(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (schoolDB *SchoolDB) UpdateStudent(student *Student) error {
	err := schoolDB.gorm.Save(student).Error
	if err != nil {
		return err
	}
	return nil
}
