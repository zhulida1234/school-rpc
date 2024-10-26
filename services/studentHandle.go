package services

import (
	"context"
	"github.com/zhulida1234/school-rpc/protobuf/student"
)

func (s *RpcServer) StudentList(ctx context.Context, request *student.StudentListRequest) (*student.StudentListResponse, error) {
	//TODO implement me
	return &student.StudentListResponse{}, nil
}

func (s *RpcServer) CreateStudent(ctx context.Context, request *student.CreateStudentRequest) (*student.CreateStudentResponse, error) {
	//TODO implement me
	return &student.CreateStudentResponse{}, nil
}

func (s *RpcServer) UpdateStudent(ctx context.Context, request *student.UpdateStudentRequest) (*student.UpdateStudentResponse, error) {
	//TODO implement me
	return &student.UpdateStudentResponse{}, nil
}
