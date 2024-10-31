package model

type StudentListRequest struct {
	PageSize uint32 `json:"pageSize"`
	PageNo   uint32 `json:"pageNo"`
}

type Student struct {
	Name      string `json:"name"`
	Age       uint32 `json:"age"`
	Gender    uint32 `json:"gender"`
	Mobile    string `json:"mobile"`
	ClassName string `json:"class"`
	Grade     uint32 `json:"grade"`
}

type StudentListResponse struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Data []Student `json:"data"`
}

type CreateStudentRequest struct {
	Name      string `json:"name"`
	Age       uint32 `json:"age"`
	Gender    uint32 `json:"gender"`
	Mobile    string `json:"mobile"`
	ClassName string `json:"className"`
	Grade     uint32 `json:"grade"`
}

type CreateStudentResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Id   uint64 `json:"id"`
}

type UpdateStudentRequest struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Age       uint32 `json:"age"`
	Gender    uint32 `json:"gender"`
	Mobile    string `json:"mobile"`
	ClassName string `json:"className"`
	Grade     uint32 `json:"grade"`
}

type UpdateStudentResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}
