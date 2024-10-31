package client

import (
	"fmt"
	"github.com/zhulida1234/school-rpc/services/model"
	"testing"
)

func TestGetStudentList(t *testing.T) {
	client := NewStudentClient("http://127.0.0.1:8970")
	result, err := client.GetStudentList(1, 10)
	if err != nil {
		fmt.Println("get support chain fail")
		return
	}
	fmt.Println("Support Chain Res:", result)
}

func TestCreateStudent(t *testing.T) {
	client := NewStudentClient("http://127.0.0.1:8970")
	req := &model.CreateStudentRequest{
		Name:      "王五",
		Age:       17,
		Mobile:    "12873263872",
		ClassName: "五三班",
		Gender:    1,
		Grade:     5,
	}

	result, err := client.CreateStudent(req)
	if err != nil {
		fmt.Println("get student list fail")
		return
	}
	fmt.Println("student list Res:", result)

}

func TestModifyStudent(t *testing.T) {
	client := NewStudentClient("http://127.0.0.1:8970")
	req := &model.UpdateStudentRequest{
		Id:        3,
		Name:      "赵六",
		Age:       17,
		Mobile:    "12873263872",
		ClassName: "五三班",
		Gender:    1,
		Grade:     5,
	}

	result, err := client.UpdateStudent(req)
	if err != nil {
		fmt.Println("update student fail")
		return
	}
	fmt.Println("update student Res:", result)

}
