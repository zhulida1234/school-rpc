package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/zhulida1234/school-rpc/services/model"
	"net/http"
)

var errWaletHTTPError = errors.New("wallet http error")

type Student struct {
	Id        uint64
	Name      string
	Age       uint32
	Gender    uint32
	Mobile    string
	ClassName string
	Grade     uint32
}

type StudentClient interface {
	GetStudentList(pageNo, pageSize uint32) (*Student, error)
}

type Client struct {
	client *resty.Client
}

func NewStudentClient(url string) *Client {
	client := resty.New()
	client.SetBaseURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode != http.StatusOK {
			method := r.Request.Method
			baseUrl := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, baseUrl, errWaletHTTPError)
		}
		fmt.Println("baseUrl::", r.Request.Method)
		fmt.Println("method::", r.Request.URL)
		fmt.Println("param::", r.Request.QueryParam)
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetStudentList(pageNo, pageSize uint32) (*model.GetStudentListResponse, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"pageNo":   fmt.Sprintf("%d", pageNo),
		"pageSize": fmt.Sprintf("%d", pageSize),
	}).SetResult(&model.GetStudentListResponse{}).Get("/api/v1/student_list")
	if err != nil {
		return nil, errors.New("get student list fail")
	}
	stl, ok := res.Result().(*model.GetStudentListResponse)
	if !ok {
		return nil, errors.New("get student list fail")
	}
	return stl, nil
}

func (c *Client) CreateStudent(request *model.CreateStudentRequest) (*model.CreateStudentResponse, error) {
	res, err := c.client.R().SetBody(request).SetResult(&model.CreateStudentResponse{}).Post("/api/v1/create_student")
	if err != nil {
		return nil, errors.New("create student fail")
	}
	ctr, ok := res.Result().(*model.CreateStudentResponse)
	if !ok {
		return nil, errors.New("create student fail")
	}
	return ctr, nil
}

func (c *Client) UpdateStudent(request *model.UpdateStudentRequest) (*model.UpdateStudentResponse, error) {
	res, err := c.client.R().SetBody(request).SetResult(&model.UpdateStudentResponse{}).Post("/api/v1/update_student")
	if err != nil {
		return nil, errors.New("update student fail")
	}
	utr, ok := res.Result().(*model.UpdateStudentResponse)
	if !ok {
		return nil, errors.New("update student fail")
	}
	return utr, nil
}
