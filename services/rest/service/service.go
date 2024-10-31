package service

import (
	"fmt"
	"github.com/zhulida1234/school-rpc/database"
	"github.com/zhulida1234/school-rpc/services/rest/model"
)

type Service interface {
	StudentList(*model.StudentListRequest) (*model.StudentListResponse, error)
	CreateStudent(*model.CreateStudentRequest) (*model.CreateStudentResponse, error)
	UpdateStudent(*model.UpdateStudentRequest) (*model.UpdateStudentResponse, error)
}

type HandleSrv struct {
	v        *Validator
	SchoolDB *database.SchoolDB
}

func NewHandlerSrv(v *Validator, schoolDB *database.SchoolDB) Service {
	return &HandleSrv{
		v:        v,
		SchoolDB: schoolDB,
	}
}

func (s *HandleSrv) StudentList(request *model.StudentListRequest) (*model.StudentListResponse, error) {
	db := s.SchoolDB
	studentList, err := db.FindStudentList(request.PageSize, request.PageNo)
	if err != nil {
		return nil, fmt.Errorf("find student list error: %v", err)
	}
	studentReturnList := make([]model.Student, len(studentList))

	for i := range studentList {
		student := model.Student{
			Name:      studentList[i].Name,
			Age:       studentList[i].Age,
			Gender:    studentList[i].Gender,
			Mobile:    studentList[i].Mobile,
			ClassName: studentList[i].ClassName,
			Grade:     studentList[i].Grade,
		}
		studentReturnList[i] = student
	}
	return &model.StudentListResponse{
		Code: 200,
		Msg:  "查询学生列表成功",
		Data: studentReturnList,
	}, nil
}

func (s *HandleSrv) CreateStudent(request *model.CreateStudentRequest) (*model.CreateStudentResponse, error) {
	db := s.SchoolDB

	studentPoint := &database.Student{
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		ClassName: request.ClassName,
		Mobile:    request.Mobile,
		Grade:     request.Grade,
	}

	res, err := db.CreateStudent(studentPoint)
	if err != nil {
		return &model.CreateStudentResponse{
			Code: 500,
			Msg:  fmt.Sprintf("create student error: %v", err),
		}, nil
	}

	return &model.CreateStudentResponse{
		Code: 200,
		Msg:  "创建学生成功",
		Id:   res.Id,
	}, nil
}

func (s *HandleSrv) UpdateStudent(request *model.UpdateStudentRequest) (*model.UpdateStudentResponse, error) {
	db := s.SchoolDB

	err := db.UpdateStudent(&database.Student{
		Id:        request.Id,
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		ClassName: request.ClassName,
		Mobile:    request.Mobile,
		Grade:     request.Grade,
	})
	if err != nil {
		return &model.UpdateStudentResponse{
			Code: 500,
			Msg:  fmt.Sprintf("update student error: %v", err),
		}, nil
	}

	return &model.UpdateStudentResponse{
		Code: 200,
		Msg:  "更新学生成功",
	}, nil

}
