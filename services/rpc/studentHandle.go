package rpc

import (
	"context"
	"fmt"
	"github.com/zhulida1234/school-rpc/database"
	"github.com/zhulida1234/school-rpc/protobuf/student"
	"strconv"
)

func (s *RpcServer) StudentList(ctx context.Context, request *student.StudentListRequest) (*student.StudentListResponse, error) {
	schoolDB := s.GetRpcSchoolDB()

	studentList, err := schoolDB.FindStudentList(request.GetPageSize(), request.GetPageNo())
	studentPointers := make([]*student.Student, len(studentList))
	for i := range studentList {
		studentPoint := &student.Student{
			Name:      studentList[i].Name,
			Age:       studentList[i].Age,
			Gender:    studentList[i].Gender,
			Mobile:    studentList[i].Mobile,
			ClassName: studentList[i].ClassName,
			Grade:     studentList[i].Grade,
		}
		studentPointers[i] = studentPoint
	}
	if err != nil {
		return nil, err
	}
	return &student.StudentListResponse{
		Code:        strconv.Itoa(200),
		Msg:         "get Student List SUCCESS",
		StudentList: studentPointers,
	}, nil
}

func (s *RpcServer) CreateStudent(ctx context.Context, request *student.CreateStudentRequest) (*student.CreateStudentResponse, error) {
	schoolDB := s.GetRpcSchoolDB()

	res, err := schoolDB.CreateStudent(&database.Student{
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		Mobile:    request.Mobile,
		ClassName: request.ClassName,
		Grade:     request.Grade,
	})
	fmt.Print(res.Id)
	if err != nil {
		return &student.CreateStudentResponse{
			Code: strconv.Itoa(500),
			Msg:  "Create Student Fail",
		}, err
	}

	return &student.CreateStudentResponse{
		Code: strconv.Itoa(200),
		Msg:  "Create Student SUCCESS",
	}, nil
}

func (s *RpcServer) UpdateStudent(ctx context.Context, request *student.UpdateStudentRequest) (*student.UpdateStudentResponse, error) {
	schoolDB := s.GetRpcSchoolDB()

	err := schoolDB.UpdateStudent(&database.Student{
		Id:        request.Id,
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		Mobile:    request.Mobile,
		ClassName: request.ClassName,
		Grade:     request.Grade,
	})

	if err != nil {
		return &student.UpdateStudentResponse{
			Code: strconv.Itoa(500),
			Msg:  "Create Student Fail",
		}, err
	}

	return &student.UpdateStudentResponse{
		Code: strconv.Itoa(200),
		Msg:  "Create Student SUCCESS",
	}, nil

}
