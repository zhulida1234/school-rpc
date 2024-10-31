package routes

import (
	"encoding/json"
	"fmt"
	"github.com/zhulida1234/school-rpc/services/rest/model"
	"net/http"
	"strconv"
)

func (h Routes) GetStudentList(w http.ResponseWriter, r *http.Request) {
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageNoStr := r.URL.Query().Get("pageNo")

	pageSize, _ := strconv.ParseUint(pageSizeStr, 10, 32)
	pageNo, _ := strconv.ParseUint(pageNoStr, 10, 32)

	cr := &model.StudentListRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	}
	stuList, err := h.svc.StudentList(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, stuList, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

func (h Routes) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var ctq model.CreateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&ctq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 调用服务方法创建学生
	crtRet, err := h.svc.CreateStudent(&ctq)
	if err != nil {
		return
	}
	err = jsonResponse(w, crtRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

func (h Routes) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var utq model.UpdateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&utq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 调用服务方法更新学生
	utrRet, err := h.svc.UpdateStudent(&utq)
	if err != nil {
		return
	}
	err = jsonResponse(w, utrRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}

}
